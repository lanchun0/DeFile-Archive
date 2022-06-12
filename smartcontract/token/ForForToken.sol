// SPDX-License-Identifier: GPL-3.0
pragma solidity  >=0.6.0 <0.9.0;
pragma experimental ABIEncoderV2;

abstract contract  ERC20Interface  {
    function totalSupply() virtual public view returns (uint);
    function balanceOf(address tokenOwner) virtual public view returns (uint balance);
    function allowance(address tokenOwner, address spender) virtual public view returns (uint remaining);
    function transfer(address to, uint tokens) virtual public returns (bool success);
    function approve(address spender, uint tokens) virtual public returns (bool success);
    function transferFrom(address from, address to, uint tokens) virtual public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract ForForToken is ERC20Interface{
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public price;
    uint256 public supply;
 
    mapping (address => uint256) public balances;
    mapping (address => mapping (address => uint256)) public allowed;

    mapping (address => uint256) public assetCount;
    mapping (address => mapping (uint256 => string)) public assetsIndex;
    mapping (address => mapping (string => Asset)) public assetsID;

    struct User {
        address addr;
        uint256 balance;
        Asset[] assets;
    }
    
    struct Asset {
        string fileID;
        uint8 permission;
    }
 
    event TopUp(address indexed tokenOwner, uint tokens);
    event Withdraw(address indexed tokenOwner, uint tokens);
    event Update(address indexed from, address indexed to, string indexed fileID, uint8 permission);

    constructor()  {
        symbol = "FOR";
        name = "ForFor Token";
        decimals = 18;
        price = 3600;
        supply = 100000000 * 10**uint256(decimals);
        // balances[msg.sender] = totalSupply;
    }
 
    function totalSupply() override public view returns (uint256) {
        return supply;
    }
 
    function balanceOf(address tokenOwner) override public view returns (uint256 balance) {
        return balances[tokenOwner];
    }
 
    function allowance(address tokenOwner, address spender) override public view returns (uint remaining) {
        return allowed[tokenOwner][spender];
    }
 
    function transfer(address to, uint256 tokens) override public returns (bool success) {
        require(to != address(0));
        require(balances[msg.sender] >= tokens);
        require(balances[to] + tokens >= balances[to]);
        balances[msg.sender] -= tokens;
        balances[to] += tokens;
        emit Transfer(msg.sender, to, tokens);
        return true;

    }
 
    function approve(address spender, uint tokens) override public returns (bool success) {
        allowed[msg.sender][spender] += tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }
 
    function transferFrom(address from, address to, uint tokens) override public returns (bool success) {
        // require(to != address(0) && from != address(0));
        require(balances[from] >= tokens);
        require(allowed[from][msg.sender] >= tokens);
        require(balances[to] + tokens >= balances[to]);
        balances[from] -= tokens;
        allowed[from][msg.sender] -= tokens;
        balances[to] += tokens;
        emit Transfer(from, to, tokens);
        return true;
    }

    function topup(uint256 _tokens) public payable returns(uint256 amount) {
        require(_tokens <= supply);
        require(_tokens > 0);
        // require(msg.value > tokens*price);
        uint256 total = _tokens*price;
        uint256 uintPrice = total / _tokens;
        uint256 remain;
        if (uintPrice != price || msg.value < total) return 0;
        if (msg.value > total) {
            remain = msg.value - total;
        } else {
            remain = 0;
        }
        payable(address(this)).transfer(total);
        if (remain != 0) {
            payable(address(msg.sender)).transfer(total);
        }
        balances[msg.sender] += _tokens;
        supply -= _tokens;
        emit TopUp(msg.sender, _tokens);
        return _tokens;
    }

    function withdraw(uint256 tokens) public returns(bool sucess){
        require(tokens > 0);
        require(balances[msg.sender] >= tokens);
        uint256 total = tokens * price;
        uint256 uintPrice = total/tokens;
        if (uintPrice != price) return false;
        uint256 balance = balances[msg.sender];
        balances[msg.sender] -= tokens;
        if (!payable(msg.sender).send(total)) {
            // balances[msg.sender] = balance;
            return false;
        }
        supply += tokens;
        emit Withdraw(msg.sender, tokens);
        return true;
    }

    function callTransferFrom(address from, address to, uint tokens)  external returns (bool success) {
        require(to != address(0) && from != address(0));
        require(balances[from] >= tokens);
        require(allowed[from][msg.sender] >= tokens);
        require(balances[to] + tokens >= balances[to]);
        balances[from] -= tokens;
        allowed[from][msg.sender] -= tokens;
        balances[to] += tokens;
        emit Transfer(from, to, tokens);
        return true;
    }

    function callApprove(address spender, uint tokens) external returns (bool) {
        allowed[msg.sender][spender] += tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }

    function updateAsset(address from, address to, string calldata _fileID, uint8 _perm) external returns (bool){
        // require(assetsID[to][_fileID].permission & _perm ==0);
        if (assetsID[to][_fileID].permission!=0){
            assetsID[to][_fileID].permission |= _perm;
        } else {
            uint count = assetCount[to];
            assetsIndex[to][count] = _fileID;
            assetCount[to]++;
            Asset storage a = assetsID[to][_fileID];
            a.fileID = _fileID;
            a.permission = _perm;
        }
        emit Update(from, to, _fileID, _perm);
        return true;
    }

    function queryUser() public view returns (User memory) {
        User memory u;
        u.addr = msg.sender;
        u.balance = balances[msg.sender];
        u.assets = new Asset[](assetCount[msg.sender]);
        for (uint256 i=0;i < assetCount[msg.sender];i++){
            string memory id = assetsIndex[msg.sender][i];
            u.assets[i].fileID = assetsID[msg.sender][id].fileID;
            u.assets[i].permission = assetsID[msg.sender][id].permission;
        }
        return u;
    }


    receive() external payable {}
}