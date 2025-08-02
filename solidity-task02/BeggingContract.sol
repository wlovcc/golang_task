// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";

contract BeggingContract is Ownable{

    //记录每个捐赠者的捐赠金额
    mapping (address => uint256) private donations;
    // 合约总捐赠金额
    uint256 private totalDonations;
    // 记录捐赠者数量
    uint256 private donorCount;
    // 记录所有捐赠者地址
    address[] private donors;

    // 捐赠开始和结束时间
    uint256 private donationStartTime;
    uint256 private donationEndTime;
    //是否启用时间限制
    bool  private timeRestrictionEnabled;


    // 事件，记录捐赠、提款、时间启用的的行为
    event Donation(address indexed donor, uint256 amount);
    event Withdrawn(address indexed owner, uint256 amount);
    event timeRestrictionUpdate(uint256 startTime,uint256 endTime,bool enable);

    // 构造函数，声明合约所有者
    constructor() Ownable(msg.sender){
        totalDonations = 0;
        donorCount = 0;

        donationStartTime = block.timestamp;
        donationEndTime = block.timestamp + 90 days;//默认90天
        timeRestrictionEnabled = false;
    }

    //允许用户向合约发送以太币，并记录捐赠信息
    function donate() public payable{
        // 捐赠数量必须大于0
        require(msg.value > 0, "Donation must be > 0");
        // 更新捐赠记录
        if(donations[msg.sender]== 0){
            donorCount ++;
            donors.push(msg.sender);
        }
        donations[msg.sender] += msg.value;
        totalDonations += msg.value;

        // 触发事件记录日志
        emit Donation(msg.sender, msg.value);
    }

    //允许合约所有者提取所有资金
    function withdraw() public  onlyOwner {
        uint256 balance = address(this).balance;
        // 合约中的资金必须大于0才能提款
        require(balance > 0, "Contract has no balance");

        // 将合约中的所有资金转移到所有者地址
        payable (owner()).transfer(balance);

        // 触发事件记录日志
        emit Withdrawn(owner(), balance);
    }

    //允许查询某个地址的捐赠金额
    function getDonation(address addr) external view returns (uint256) {
        return donations[addr];
    }

    // 查询合约余额
    function getBalance() public view returns(uint) {
        return address(this).balance;
    }

    // 查询捐赠者总数
    function getDonorCount() public view returns (uint256) {
        return donorCount;
    }

    // 查询所有捐赠者地址
    function getAllDonors() public view returns (address[] memory) {
        return donors;
    }

    // 回退函数
    receive() external payable {
        donate();
    }

    fallback() external payable {
        donate();
    }

    //获取捐赠排行榜前N名
    function getTopDonorssort(uint256 topN) public view returns (address[] memory addr, uint256[] memory count) {
        require(topN > 0 && topN <=  donorCount, "Invalid topN value");
        address[] memory tmpDoors = new address[](donorCount);
        uint256[] memory tmpCount = new uint256[](donorCount);

        for(uint256 i=0; i<donorCount; i++){
            tmpDoors[i] = donors[i];
            tmpCount[i] = donations[donors[i]];
        }

        for(uint256 i=0; i<donorCount; i++){
            for(uint256 j=0; j<donorCount-i-1; j++){
                if(tmpCount[j] < tmpCount[j+1]){
                    // 交换金额
                    uint256 tmpC = tmpCount[j];
                    tmpCount[j] = tmpCount[j+1];
                    tmpCount[j+1] = tmpC;

                    //交换地址
                    address tmpAddr = tmpDoors[j];
                    tmpDoors[j] = tmpDoors[j+1];
                    tmpDoors[j+1] = tmpAddr;
                }
            }
        }

        addr = new address[](topN);
        count = new uint256[](topN);
        for (uint256 i=0; i<topN; i++) {
           addr[i] = tmpDoors[i];
           count[i] = tmpCount[i];
        }
        
        return (addr,count);
    }


    //设置捐赠时间限制
     function setTimeRestriction(uint256 startTime,uint256 endTime,bool enable) public onlyOwner {
        require(startTime < endTime,"End time must be after start time");
        donationStartTime = startTime;
        donationEndTime = endTime;
        timeRestrictionEnabled = enable;
     }

}

// https://sepolia.etherscan.io/address/0x33ae2190949e4637f3a6a2134fb393467fd42b34