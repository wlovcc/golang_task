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
    
