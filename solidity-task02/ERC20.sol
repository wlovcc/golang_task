// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
/*
interface IERC20Metadata {
    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 value);
    // 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);
    //返回当前代币的总供应量
    function totalSupply() external view returns (uint256);
    //查询账户余额
    function balanceOf(address account) external view returns (uint256);
    //转账
    function transfer(address to, uint256 value) external returns (bool);
    //查询授权机制中的剩余额度
    function allowance(address owner, address spender) external view returns (uint256);
    //授权
    function approve(address spender, uint256 value) external returns (bool);
    //代扣转账
    function transferFrom(address from, address to, uint256 value) external returns (bool);
    
}
*/
/**
* 基于ERC20标准的代币合约
*/
//contract ERC20_Task is IERC20Metadata { 
contract ERC20_Task { 

    string private _name;   //代币(合约)名称
    string private _symbol; //代币(合约)符号

    uint8  _decimals;   //小数位数
    uint256 private _totalSupply;   //供应量

    address public _owner;   //合约所有者

    mapping(address => uint256) balances;  //余额映射
    mapping(address => mapping(address => uint256)) approves; //授权映射：owner => spender => amount


    // constructor(){
    //     _name = "cc_mytokens";              //代币(合约)名称
    //     _symbol = "CTC_TK";                 //代币(合约)简称
    //     _owner=msg.sender;                  //代币(合约)地址
    //     _decimals =26;                      //小数位
    //     _totalSupply = 100*10**9 ;          //发行总代币
    //     balances[_owner] = 100 * 10 ** 9;   //余额

    //     emit Transfer(address(0), _owner, _totalSupply);  
    // }

    // 转账事件
    event Transfer(address indexed from, address indexed to, uint256 amount);
    // 授权事件
    event Approval(address indexed owner, address indexed spender, uint256 value);
    // 定义 onlyOwner修饰器
    constructor(string memory _name_,string memory _symbol_,uint8 _decimals_,uint256 _totalSupply_) {
        _name = _name_;
        _symbol = _symbol_;
        _decimals = _decimals_;
        _owner = msg.sender;
        //_decimals =26;                      //小数位
        _totalSupply = _totalSupply_ ;          //发行总代币
        balances[_owner] = _totalSupply_;   //余额

        emit Transfer(address(0), _owner, _totalSupply);
    }

    //返回当前代币的总供应量
    function totalSupply() external view returns (uint256){
        return _totalSupply;
    }

    //查询账户余额
    function balanceOf(address account) external view returns (uint256){
        return balances[account];
        //return account.balance;
    }

    //转账-转移代币:  toAddress：代币接收者，仅被动接收代币，无需任何权限
    function transfer(address to, uint256 value) external returns (bool){
        require(to != address(0),"ExampleERC20:  not transfer to zero address");
        require(balances[_owner] >= value,"ERC20: balance is not enough");

        // 转移代币
        balances[_owner] -= value;        // 1。减少转出账户余额
        balances[to] += value;          // 2.增加转入账户余额

        emit Transfer(_owner, to, value); // 触发转账事件
        return true;

    }
    //查询授权机制中的剩余额度
    function allowance(address owner, address spender) external view returns (uint256){
        return approves[owner][spender];
    }

    //授权  spender(被授权者)
    function approve(address spender, uint256 value) external returns (bool){
        require(_owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0),"ERC20: approve to the zero address");
        require(approves[_owner][spender] == 0,"ERC20: Spender is approved");
        require(approves[_owner][spender] >= value,"no enough allowance");

        approves[_owner][spender] = value;
        emit Approval(_owner, spender, value);
        return true;
    }
    
    //代扣转账 from就是spender
    function transferFrom(address from, address to, uint256 value) public virtual returns (bool) {
        //1. _transfer(from, to, amount)：转移代币，确保 from 有足够余额
        //2. _spendAllowance(from, _msgSender(), amount)：扣减授权额度（或验证无限授权）
        //address spender = msg.sender;   //_owner
        require(balances[from] >= value,"transferFrom: not enough balance");
        uint256 currentAllowance = approves[from][_owner];
        require(currentAllowance >= value,"ERC20: insufficient allowance");

        // 先转账
        balances[from] -= value;        // 1。减少转出账户余额
        balances[to] += value;          // 2.增加转入账户余额
        // 后扣减额度
        approves[from][_owner] -= value; //3.减少from对msg.sender(调用者)的授权额度

        emit Transfer(from, to, value); // 触发转账事件
        return true;
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, "Only owner can call this function");
        _;
    }

    // 增发代币
    function mint(address to, uint256 _value) public onlyOwner {
        require(to != address(0),"mint not to zero address");
        uint256 amount = _value * 10 **uint256(_decimals);
        _totalSupply += amount;
        balances[to] += amount;
        //代币从 零地址（address (0)） 转移到目标地址 _toAddress，实际上意味着这些代币是新创建的（增发）
        emit Transfer(address(0), to, amount);
    }//
}