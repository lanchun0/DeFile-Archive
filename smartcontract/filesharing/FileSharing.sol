// SPDX-License-Identifier: GPL-3.0
pragma solidity   >=0.6.0 <0.9.0;
pragma experimental ABIEncoderV2;

    struct File {
        // string fileID;
        address owner;
        Metedata meteData;
        uint8 permissionLevel;
        bool tradable;
        uint256 price;
        uint reward;
        mapping (address=>uint8) permissionList; // address => permission
        mapping (uint256=>address) listIndex;
        uint256 listCount;
    }
    struct Metedata {
        string hashDigest;
        address signer;
        string fileName;
        uint size;
        string timestamp;
    }
    struct PermissionList {
        address user;
        uint8 permission;
    }
    
    struct QueryFile {
        string fileID;
        address owner;
        Metedata meteData;
        bool tradable;
        uint256 price;
        uint8 permissionLevel;
        PermissionList[] permissionList;   
    }

contract FileSharing {
    event CreateFile(address indexed _owner, string indexed _fileID, uint8 _pL);
    event WriteFile(address indexed _writer, string indexed _fileID,  string indexed _digest);
    event Transfer(address indexed _from, address indexed _to, string indexed _fileID, uint8 _permission);

    // address behavior;
    mapping (string => File) files;   // fileID => File
    mapping (string => string) id2IPFS; //  fileID => IPFS cid
    mapping (string => uint8) permLevels;
    mapping (string => uint8) perms;
    mapping (uint256 => string) index; // index => fileid
    uint256 fileAmount;   // total number of file

    address token;

    constructor(address _addr)  {
        // behavior = _behavior;
        permLevels["L_0"] = 0x80;
        permLevels["L_1"] = 0x40;
        permLevels["L_2"] = 0x20;
        permLevels["L_3"] = 0x10;
        permLevels["L_4"] = 0x08;
        perms["owner"] = 0x80;
        perms["admin"] = 0x40;
        perms["writer"] = 0x20;
        perms["reader"] = 0x10;
        token = _addr;
     }

    // constructor () public {
    //     fileAmount = 0;
    // }

    function createFile(string calldata _fileID, uint8 _pL, bool _tradable, uint256 _price) public returns(uint256) {
        require(files[_fileID].owner==address(0),"Token already exists");
        require((_pL==0x80||_pL==0x40||_pL==0x20||_pL==0x10||_pL==0x08),"invalid permission level");
        require((_tradable&&(_pL<=0x40))||!_tradable,"only permissioned file can be traded");
        require(ForForToken(token).updateAsset(msg.sender,  msg.sender, _fileID, 0xf0),"failed to update the asset");
        File storage f = files[_fileID];
        f.owner = msg.sender;
        f.permissionLevel = _pL;
        f.tradable = _tradable;
        f.price = _price;
        f.permissionList[msg.sender] = perms["owner"] | perms["admin"] | perms["writer"] | perms["reader"];
        f.listIndex[0] = msg.sender;
        f.listCount = 1;
        index[fileAmount] = _fileID;
        fileAmount++;
        emit CreateFile(msg.sender,  _fileID,  _pL);
        return fileAmount;
    } 

    function writeFile(string calldata _fileID, string calldata _digest, string calldata _fileName, uint _size, string calldata _time, string calldata _ipfsHash)public {
        require((files[_fileID].permissionList[msg.sender]&perms["writer"])!=0,"No authority to write");
        require(files[_fileID].permissionLevel<permLevels["L_0"]||files[_fileID].permissionList[msg.sender]&perms["owner"]!=0,"No autority to write");
        files[_fileID].meteData.hashDigest = _digest;
        files[_fileID].meteData.fileName = _fileName;
        files[_fileID].meteData.size = _size;
        files[_fileID].meteData.timestamp = _time;
        files[_fileID].meteData.signer = msg.sender;
        id2IPFS[_fileID] = _ipfsHash;
        emit WriteFile(msg.sender,  _fileID,  _digest);
    }

    function readFile(string calldata _fileID) public view returns(string memory ipfsHash, QueryFile memory) {
        require(files[_fileID].permissionList[msg.sender]>=perms["owner"]||files[_fileID].permissionLevel<permLevels["L_0"],"No autority to read");
        require(files[_fileID].permissionList[msg.sender]>=perms["reader"]||files[_fileID].permissionLevel<=permLevels["L_3"],"No autority to read");
        QueryFile memory f;
            // File memory f = files[_fileID]
        f.fileID = _fileID;
        f.price = files[_fileID].price;
        f.tradable = files[_fileID].tradable;
        f.owner = files[_fileID].owner;
        f.permissionLevel = files[_fileID].permissionLevel;
        f.meteData.hashDigest = files[_fileID].meteData.hashDigest;
        f.meteData.signer = files[_fileID].meteData.signer;
        f.meteData.fileName = files[_fileID].meteData.fileName;
        f.meteData.size = files[_fileID].meteData.size;
        f.meteData.timestamp = files[_fileID].meteData.timestamp;
        f.permissionList = new PermissionList[](files[_fileID].listCount);
        for (uint256 i=0;i < files[_fileID].listCount;i++){
            // uint8 perm = files[pub].permissionList[pub];
            f.permissionList[i].user = files[_fileID].listIndex[i];     
            f.permissionList[i].permission = files[_fileID].permissionList[f.permissionList[i].user];
        }
        return (id2IPFS[_fileID],f);
    } 

    function authorize(address _to, string calldata _fileID, uint8 _perm) public  {
        require(_perm==perms["admin"]||_perm==perms["writer"]||_perm==perms["reader"],"invalide permission");
        require(files[_fileID].permissionLevel<permLevels["L_0"],"private file cannot be authorized");
        require(files[_fileID].permissionList[msg.sender]>=perms["admin"],"only owner or admin can authorize");
        require(_perm!=perms["admin"]||files[_fileID].permissionList[msg.sender]>=perms["owner"],"permission denied");
        require((files[_fileID].permissionLevel!=permLevels["L_1"]&&files[_fileID].permissionLevel!=permLevels["L_3"])||_perm!=perms["admin"],"permission denied");
        require(files[_fileID].permissionList[msg.sender]>=perms["admin"],"permission denied");
        require(ForForToken(token).updateAsset(msg.sender,  _to,  _fileID, _perm),"failed to update the asset");
        uint8 p = files[_fileID].permissionList[_to];
        if (p==0) {
            uint256 count = files[_fileID].listCount;
            files[_fileID].listIndex[count] = _to;
            files[_fileID].listCount++;  
        }
        if (_perm==perms["admin"]) p = p | _perm | perms["writer"] | perms["reader"];
        if (_perm==perms["writer"]) p = p| _perm | perms["reader"];
        if (_perm==perms["reader"]) p = p| _perm;
        files[_fileID].permissionList[_to] = p;
        emit Transfer(msg.sender,  _to, _fileID,  _perm);
    }
    
    function queryFile(string calldata _fileID) public view returns(QueryFile memory) {
        require(files[_fileID].permissionList[msg.sender]!=0x00||files[_fileID].permissionLevel<=permLevels["L_3"]||files[_fileID].tradable,"no autority");
        QueryFile memory f;
        f.fileID = _fileID;
        f.owner = files[_fileID].owner;
        f.price = files[_fileID].price;
        f.tradable = files[_fileID].tradable;
        f.permissionLevel = files[_fileID].permissionLevel;
        f.meteData.hashDigest = files[_fileID].meteData.hashDigest;
        f.meteData.signer = files[_fileID].meteData.signer;
        f.meteData.fileName = files[_fileID].meteData.fileName;
        f.meteData.size = files[_fileID].meteData.size;
        f.meteData.timestamp = files[_fileID].meteData.timestamp;
        f.permissionList = new PermissionList[](files[_fileID].listCount);
        for (uint256 i=0;i < files[_fileID].listCount;i++){
            f.permissionList[i].user = files[_fileID].listIndex[i];     
            f.permissionList[i].permission = files[_fileID].permissionList[f.permissionList[i].user];
        }
        return f;
    }

    function queryFileByIndex(uint256 _index) public view returns (QueryFile memory) {
        require(_index < fileAmount,"File does not exist");
        string memory _fileID = index[_index];
        require(files[_fileID].permissionList[msg.sender]!=0x00||files[_fileID].permissionLevel<=permLevels["L_3"]||files[_fileID].tradable,"no autority");
        QueryFile memory f;
        f.fileID = _fileID;
        f.owner = files[_fileID].owner;
        f.price = files[_fileID].price;
        f.tradable = files[_fileID].tradable;
        f.permissionLevel = files[_fileID].permissionLevel;
        f.meteData.hashDigest = files[_fileID].meteData.hashDigest;
        f.meteData.signer = files[_fileID].meteData.signer;
        f.meteData.fileName = files[_fileID].meteData.fileName;
        f.meteData.size = files[_fileID].meteData.size;
        f.meteData.timestamp = files[_fileID].meteData.timestamp;
        f.permissionList = new PermissionList[](files[_fileID].listCount);
        for (uint256 i=0;i < files[_fileID].listCount;i++){
            f.permissionList[i].user = files[_fileID].listIndex[i];     
            f.permissionList[i].permission = files[_fileID].permissionList[f.permissionList[i].user];
        }
        return f;
    }

    function getFilePrice(string calldata _fileID) public view returns (uint) {
        require(files[_fileID].tradable,"file cannot be traded");
        return files[_fileID].price;
    }

    function tradeFile(string calldata _fileID) public returns (bool) {
        require(files[_fileID].tradable,"file cannot be traded");
        bool success;
        uint price = files[_fileID].price;
        // bytes memory transfer = abi.encodeWithSignature("transferFrom(address, address, uint)",msg.sender, address(this), price);
        // bytes memory transfer_owner = abi.encodeWithSignature("approve(address, uint)" ,address(files[_fileID].owner), (price/10)*2);
        // bytes memory transfer_signer = abi.encodeWithSignature("approve(address, uint)", address(files[_fileID].meteData.signer), (price/10)*8);
        success = ForForToken(token).callTransferFrom(msg.sender, address(this), price);
        // (success,) = bank.call(transfer);
        if (success) {
            if (files[_fileID].permissionList[msg.sender]==0) {
                files[_fileID].listIndex[files[_fileID].listCount] = msg.sender;
                files[_fileID].listCount++;
            }
            files[_fileID].permissionList[msg.sender] |= perms["reader"];
            if (files[_fileID].permissionLevel<=permLevels["L_3"]) {
                files[_fileID].permissionList[msg.sender] |= perms["writer"];
            } 
            require(ForForToken(token).updateAsset(address(this),  msg.sender,  _fileID, files[_fileID].permissionList[msg.sender]));
        } else {
            return false;
        }
        uint to_signer = (price*80)/100;
        uint to_owner = price - to_signer;
        if (price < 10 || files[_fileID].meteData.signer==address(0)) {
            return ForForToken(token).callApprove(files[_fileID].owner, price);
        }
        
        success = ForForToken(token).callApprove(files[_fileID].owner, to_owner);
        if (!success) {
            return false;
        }
        success = ForForToken(token).callApprove(files[_fileID].meteData.signer, to_signer);
        return success;
    }

    function getFileCount() public view returns (uint256) {
        return fileAmount;
    }

}
 
interface ForForToken {
    function callTransferFrom(address from, address to, uint tokens)  external returns (bool success);
    function callApprove(address spender, uint tokens) external returns (bool);
    function updateAsset(address from, address to, string calldata _fileID, uint8 _perm) external returns (bool);
}
