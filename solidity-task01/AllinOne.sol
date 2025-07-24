//evm 强类型脚本语言
// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 < 0.9.0;

import "hardhat/console.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// 投票合约，继承 Ownable 实现权限控制
contract Voting is Ownable{
    // 候选人得票数映射
    mapping (string => uint ) public votes;
    // 候选人有效性映射
    mapping(string => bool) public isCandidate;
    // 地址投票状态
    mapping(address => bool) public hasVotedState;
    // 候选人列表
    string[] public votersArr;
    // 投票者地址列表
    address[] public persionArr;

    // 构造函数：设置部署者为所有者
     constructor() Ownable(msg.sender) {
        romanMap['I'] = 1;
        romanMap['V'] = 5;
        romanMap['X'] = 10;
        romanMap['L'] = 50;
        romanMap['C'] = 100;
        romanMap['D'] = 500;
        romanMap['M'] = 1000;
     }

    // 允许用户投票给某个候选人
    function vote(string memory user) public  {
        //异常判断
        require(bytes(user).length > 0,unicode"候选人名称不能为空");
        require(!hasVotedState[msg.sender], unicode"已投票");

        if(votes[user] == 0){
            votersArr.push(user);
            isCandidate[user] = true;
        }
        votes[user]++;
        hasVotedState[msg.sender] = true;
        persionArr.push(msg.sender);
        
    }

    // 返回某个候选人的得票数
    function getVotes(string memory user) public view returns (uint){
        require(bytes(user).length > 0,unicode"候选人名称不能为空");
        return votes[user];
    }

    //重置所有候选人的得票数
    function resetVotes() external onlyOwner{
        // 重置候选人票数
        for(uint i=0;i <votersArr.length;i++){
            votes[votersArr[i]] = 0;
        }
        // 重置投票人状态
        for (uint i=0; i<persionArr.length;i++){
            hasVotedState[persionArr[i]] = false;
        }
        // 清空投票者列表
        delete persionArr;
    }

    //查看所有候选人的得票数
    function allVotes() public view {
        for (uint i = 0; i < votersArr.length; i++) {
            console.log("Candidate: %s, Votes: %d", votersArr[i], votes[votersArr[i]]);
        }
    }

    // 2. 反转字符串
    function reverseString(string memory strIn) public pure returns (string memory) {
        bytes memory str = bytes(strIn);
        uint len = str.length;

        for(uint i=0; i<len/2; i++){
            bytes1 tmp = str[i];
            str[i] = str[len-i-1];
            str[len-i-1] = tmp;
        }
        return string(str);
    }

    // 3.罗马数字转整数
    mapping(bytes1 => uint) public romanMap ;   // 数据再构造函数中赋值
    function romanToInt(string memory strIn) public view returns (uint) {
        uint v;
        uint lv;
        uint cv;
        uint len = bytes(strIn).length;
        bytes memory str = bytes(strIn);

        for (uint i=len; i > 0; i--) {
            cv = romanMap[(str[i-1])];
            if (cv < lv) {
                v -= cv;
            } else {
                v += cv;
            }
            lv = cv;
        }
        return v;
    }

    // 4.整数转罗马数字
    function intToRoman(uint num) public pure returns  (string memory){
        require(num > 0 && num < 4000, "Number must be between 1 and 3999 for Roman numeral conversion.");

        string[10][4] memory R;
        R[0] = ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"];
        R[1] = ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"];
        R[2] = ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"];
        R[3] = ["", "M", "MM", "MMM", "", "", "", "", "", ""];

        return string(abi.encodePacked(
            R[3][num / 1000],    // 千位
            R[2][(num % 1000) / 100], // 百位
            R[1][(num % 100) / 10],   // 十位
            R[0][num % 10]            // 个位
        ));
    }

    function merge(uint256[] memory arr1, uint256[] memory arr2) public pure returns (uint256[] memory) {
        uint256[] memory result = new uint256[](arr1.length + arr2.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        while(i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                result[k++] = arr1[i++];
            } else {
                result[k++] = arr2[j++];
            }
        }

        while(i < arr1.length) {
            result[k++] = arr1[i++];
        }

        while(j < arr2.length) {
            result[k++] = arr2[j++];
        }
        return result;
    }

    function binarySearch(uint256[] memory arr, uint256 target) public pure returns (int256) {
        require(arr.length == 0 , unicode"数组不能为空");

        int256 left = 0;
        int256 right = int256(arr.length)-1;

        while(left <= right){
            int256 mid = (left+right)-1;
            if(arr[uint256(mid)] == target){
                return mid;
            }else if(arr[uint256(mid)] < target){
                left = mid+1;
            }else{
                right = mid-1;
            }
        }
        return -1;
    }
}
