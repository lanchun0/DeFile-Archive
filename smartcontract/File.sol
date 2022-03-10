// SPDX-License-Identifier: GPL-3.0
pragma solidity  >=0.6.0 <0.9.0;
pragma experimental ABIEncoderV2;

    struct File {
        string fileID;
        string hashDigest;
        string owner;
        bool isUsed;
        Metedata meteData;
        uint8 permissionLevel;
        mapping (string=>uint8) permissionList; // publickey => permission
        mapping (uint256=>string) listIndex;
        uint256 listCount;
        // PermissionList[] permissionList;
        string signature;
    }

    struct Metedata {
        string fileName;
        uint size;
        string timestamp;
    }

    struct OffFile {
        string fileID;
        string hashDigest;
        string owner;
        Metedata meteData;
        uint8 permissionLevel;
        PermissionList[] permissionList;
        string signature;
    }

    struct PermissionList {
        string publicKey;
        uint8 permission;
    }

contract FileSharing {
    event CreateFile(string indexed _fileID, string indexed _owner, uint8 _pL);
    event WriteFile(string indexed _fileID, string indexed _writer,  string _digest);
    event ReadFile(string indexed _fileID, string indexed _reader);

    address behavior;
    mapping (string => uint8) permLevels;
    mapping (string => uint8) perms;
    mapping (uint64 => string) keys;
    mapping (string => File) files;
    uint64 fileAmount;

    constructor(address _behavior) public {
        behavior = _behavior;
        permLevels["L_0"] = 0x80;
        permLevels["L_1"] = 0x40;
        permLevels["L_2"] = 0x20;
        permLevels["L_3"] = 0x10;
        permLevels["L_4"] = 0x08;
        perms["owner"] = 0x80;
        perms["admin"] = 0x40;
        perms["writer"] = 0x20;
        perms["reader"] = 0x10;
     }

    // constructor () public {
    //     fileAmount = 0;
    // }

    function createFile(string calldata _fileID, string calldata _digest, string calldata _owner, uint8 _pL) public returns(uint64) {
        require(!existFile(_fileID),"File already exists");
        require(abi.encodePacked(_fileID).length!=0,"file id cannot be nil");
        require((_pL==0x80||_pL==0x40||_pL==0x20||_pL==0x10||_pL==0x08),"invalid permission level");
        bool added = false;
        uint8 perm = perms["owner"] | perms["admin"] | perms["writer"] | perms["reader"];
        // add file into behavior contract initially
        added = Behavior(behavior).addFile(_owner, _fileID, perm);
        if (!added) return fileAmount;
        fileAmount++;
        keys[fileAmount] = _fileID;
        File storage f = files[_fileID];
        f.fileID = _fileID;
        f.isUsed = true;
        f.hashDigest = _digest;
        f.owner = _owner;
        f.permissionLevel = _pL;
        f.permissionList[_owner] = perm;
        f.listCount = 1;
        f.listIndex[0] = _owner;
        emit CreateFile(_fileID, _owner, _pL);
        return fileAmount;
    }

    function writeFile(string calldata _pub, string calldata _fileID, string calldata _digest, string calldata _signature) public returns(bool){
        if(!canWrite(_pub, _fileID)) return false;
        files[_fileID].hashDigest = _digest;
        files[_fileID].signature = _signature;
        emit WriteFile(_fileID, _pub , _digest);
        return true;
    }

    function writeMeta(string calldata _pub, string calldata _fileID, string calldata _fileName, uint _size, string calldata _time) public returns(bool){
        if(!canWrite(_pub, _fileID)) return false;
        // files[_fileID].fileID = _fileID;
        files[_fileID].meteData.fileName = _fileName;
        files[_fileID].meteData.size = _size;
        files[_fileID].meteData.timestamp = _time;
        return true;
    }


    // function writeFile(string calldata _pub, OffFile memory _oF) public returns(bool){
    //     string memory _fileID = _oF.fileID;
    //     if(!canWrite(_pub, _fileID)) return false;
    //     files[_fileID].hashDigest = _oF.hashDigest;
    //     files[_fileID].signature = _oF.signature;
    //     files[_fileID].meteData.fileName = _oF.meteData.fileName;
    //     files[_fileID].meteData.size = _oF.meteData.size;
    //     files[_fileID].meteData.timestamp = _oF.meteData.timestamp;
    //     emit WriteFile(_oF.fileID,_pub , _oF.hashDigest);
    //     return true;
    // }

    function readFile(string calldata _pub, string calldata _fileID) public returns(OffFile memory){
        OffFile memory f;
        if(!canRead(_pub, _fileID)) return f;
        emit ReadFile(_fileID,_pub);
        return queryFile(_fileID);
    }

    function addPermission(string calldata _from, string calldata _to, string calldata _fileID, uint8 _perm) public returns(bool) {
        if(!existFile(_fileID)) return false;
        if(!Behavior(behavior).existBehavior(_to)) return false;
        uint8 perm = 0;
        bool canAdd = false;
        if(_perm == perms["admin"] && _isOwner(_from,_fileID)) {
            perm = _getPermission(_to, _fileID);
            perm = perm|perms["admin"]|perms["writer"]|perms["reader"];       
            canAdd = true;
            // files[_fileID].PermissionList[_to] = perm;
            // return true;
        }
        if (_isAdmin(_from, _fileID)) {
            if(_perm == perms["writer"]) {
                perm = _getPermission(_to, _fileID);
                perm = perm|perms["writer"]|perms["reader"];
                canAdd = true;
            }
            if(_perm == perms["reader"]) {
                perm = _getPermission(_to, _fileID);
                perm = perm|perms["reader"];
                canAdd = true;
            }
        }
        if(!canAdd) return false;
        if (!Behavior(behavior).authorFile(_from, _to, _fileID, perm)) return false;
        if(files[_fileID].permissionList[_to]==0) {
            files[_fileID].listCount++;
            files[_fileID].listIndex[files[_fileID].listCount-1] = _to;
        }
        files[_fileID].permissionList[_to] = perm;
        return true;
    }


    function _isOwner(string memory _pub, string memory _fileID) private view returns(bool) {
        if(files[_fileID].permissionList[_pub]&perms["owner"]!=0) return true;
        return false;
    }

    function _isAdmin(string calldata _pub, string calldata _fileID) private view returns(bool) {
        if(files[_fileID].permissionList[_pub]&perms["admin"]!=0) return true;
        return false;
    }


    function queryFile(string memory _fileID) public view returns(OffFile memory){
        // require(isExistFile(_fileID),"File do not exist");
        File storage f = files[_fileID];
        OffFile memory oF;
        oF.fileID = f.fileID;
        oF.hashDigest = f.hashDigest;
        oF.owner = f.owner;
        oF.permissionLevel = f.permissionLevel;
        oF.signature = f.signature;
        oF.meteData.fileName = f.meteData.fileName;
        oF.meteData.size = f.meteData.size;
        oF.meteData.timestamp = f.meteData.timestamp;
        oF.permissionList = new PermissionList[](f.listCount);
        for(uint256 i=0;i < f.listCount;i++){
            string memory pub = files[_fileID].listIndex[i];
            // uint8 perm = files[pub].permissionList[pub];
            oF.permissionList[i].publicKey = pub;     
            oF.permissionList[i].permission = _getPermission(pub, _fileID);
        }
        return oF;
    }

    function queryFileByIndex(uint64 _index) public view returns (OffFile memory) {
        // require(_index <= fileAmount,"File do not exist");
        OffFile memory oF;
        string memory key = keys[_index];
        oF = queryFile(key);
        return oF;
    }
    

    function getFileCount() public view returns (uint64) {
        return fileAmount;
    }

    function existFile(string memory _fileID) public view returns(bool){
        return files[_fileID].isUsed;
    }

    function canWrite(string memory _writer, string memory _fileID) public view returns(bool){
        if(!existFile(_fileID)) return false;
        if(_isOwner(_writer, _fileID)) return true;
        uint8 pL = files[_fileID].permissionLevel;
        if (pL==permLevels["L_0"]) return false;
        if(files[_fileID].permissionList[_writer] & perms["writer"] != 0) return true;
        return false;
    }

    function canRead(string calldata _reader, string calldata _fileID) public view returns(bool) {
        if(!existFile(_fileID)) return false;
        if(_isOwner(_reader, _fileID)) return true;
        uint8 pL = files[_fileID].permissionLevel;
        if (pL==permLevels["L_0"]) return false;
        if (pL==permLevels["L_3"]||pL==permLevels["L_4"]) return true; //public
        if(files[_fileID].permissionList[_reader] & perms["reader"] !=0) return true;
        return false;
    }
    
    function _getPermission(string memory _pub, string memory _fileID) private view returns(uint8) {
        return files[_fileID].permissionList[_pub];
    }

    function getUserCount()public view returns(uint64) {
        return Behavior(behavior).getUserCount();
    }
}
 
interface Behavior {
    function getUserCount() external view returns(uint64);
    function existBehavior(string calldata _pub) external view returns(bool);
    function hasFile(string calldata  _pub, string calldata _fileID) external view returns(bool, uint8);
    function addFile(string calldata _pub, string calldata  _fileID, uint8 _permission) external returns(bool);
    function authorFile(string calldata _from, string calldata _to, string calldata _fileID, uint8 _permission) external returns(bool);
}



 