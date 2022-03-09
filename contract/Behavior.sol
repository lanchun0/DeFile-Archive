// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.9.0; 
pragma experimental ABIEncoderV2;

struct User {
    string publicKey;
    string privateKey;
    Asset[] assets;
    bool exist;
}

struct Asset {
    string fileid;
    uint8 permission;
}


contract Behavior {
    mapping (uint64 => string) keys;
    mapping (string => User) users;
    uint64 userAmount;

    function createUser(string calldata _pub, string calldata _priv) public returns(uint64) {
        require(abi.encodePacked(_pub).length!=0,"public key cannot be nil");
        require(abi.encodePacked(_priv).length!=0,"private key cannot be nil");
        require(!existUser(_pub),"user already exists");
        userAmount++;
        keys[userAmount] = _pub;
        User storage u = users[_pub];
        u.privateKey = _priv;
        u.exist = true;
        return userAmount;
    }

    function existUser(string calldata _pub) public view returns(bool){
        return users[_pub].exist;
    }

    function getUserCount() external view returns(uint64){
        return userAmount;
    }

    function getUser(string calldata _pub) public view returns(User memory) {
        require(existUser(_pub),"user does not exist");
        User memory u = users[_pub];
        return u;
    }

    function hasAsset(string calldata  _pub, string calldata _fileID) public view returns(bool, uint8){
        if (bytes(_fileID).length==0 || !existUser(_pub)) {
            return (false, 0);
        }
        User memory u = getUser(_pub);
        bool has=false;
        uint8 permission=0;
        for (uint i=0; i < u.assets.length; i++){
            string memory fid = u.assets[i].fileid;
            if(keccak256(abi.encodePacked(fid)) == keccak256(abi.encodePacked(_fileID))) {
                has = true;
                permission = u.assets[i].permission;
                break;
            }
        }
        return (has, permission);
    } 

    function addAsset(string calldata _pub, string calldata  _fileID, uint8 _permission) public returns(bool) {
        if (!existUser(_pub) || bytes(_fileID).length==0) return false;
        // if(_permission != 0x80 && _permission != 0x40 && _permission != 0x20 && _permission != 0x10){
        //     return false;
        // }
        bool has = false;
        (has, ) = hasAsset(_pub, _fileID);
        if (has) {
           return false;
        }
        users[_pub].assets.push();
        uint256 newIndex = users[_pub].assets.length - 1;
        users[_pub].assets[newIndex].fileid = _fileID;
        users[_pub].assets[newIndex].permission = _permission;
        return true;
    }

    function transferAsset(string calldata _from, string calldata _to, string calldata _fileID, uint8 _permission) public returns(bool){
        // require((_permission==0x40|| _permission==0x20||_permission==0x10),"illegal permission");
        bool hasF=false;
        bool hasT=false;
        uint8 permF=0;
        uint8 permT=0;
        (hasF, permF) = hasAsset(_from, _fileID);
        (hasT, permT) = hasAsset(_to, _fileID);
        if (!existUser(_to) || !hasF || permF < _permission || permF < 0x40 || permT >= _permission) {
            return false;
        }
        if (!hasT) {
            if(!addAsset(_to, _fileID, _permission)){
                return false;
            }
            return true;
        }
        for (uint i=0; i < users[_to].assets.length; i++) {
            string memory fid = users[_to].assets[i].fileid;
            if (keccak256(abi.encodePacked(fid)) == keccak256(abi.encodePacked(_fileID))){
                users[_to].assets[i].permission = _permission | permT;
                return true;
            }
        }
        return false;
    }
    
    function existBehavior(string calldata _pub) external view returns(bool) {
        return existUser(_pub);
    }

    function hasFile(string calldata  _pub, string calldata _fileID) external view returns(bool, uint8){
        return hasAsset(_pub, _fileID);
    }

    function addFile(string calldata _pub, string calldata  _fileID, uint8 _permission) external returns(bool){
        return addAsset(_pub, _fileID, _permission);
    }

    function authorFile(string calldata _from, string calldata _to, string calldata _fileID, uint8 _permission) external returns(bool){
        return transferAsset(_from, _to, _fileID, _permission);
    }
}