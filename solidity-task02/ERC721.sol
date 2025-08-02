// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";         // 权限控制：仅所有者可执行敏感操作
//import "@openzeppelin/contracts/token/ERC721/ERC721.sol";   // ERC721标准实现
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol"; // 支持每个代币独立URI

contract  ERC721_Task is ERC721URIStorage, Ownable {

    constructor() ERC721("ctc_task721", "CTC") Ownable(msg.sender) {}

    uint256 private _nextTokenId;   // 状态变量：下一个可用的tokenId（自动递增）

    /**
     * @dev 铸造新NFT并关联元数据（仅合约所有者可调用）
     * @param recipient NFT接收者地址
     * @param tokenURI 元数据IPFS链接
     */
    function mintNFT(address recipient, string memory tokenURI) public onlyOwner {
        uint256 tokenId = _nextTokenId;       // 获取当前tokenId
        _nextTokenId++;                       // 递增计数器，为下次铸造准备
        _safeMint(recipient, tokenId);               // 安全铸造（检查接收方是否为合约）
        _setTokenURI(tokenId, tokenURI);           // 关联元数据URI到tokenId
    }

    mapping(uint256 => string) private _tokenURIs;
    function getTokenURI(uint256 tokenId) public view returns (string memory) {
        return _tokenURIs[tokenId];
    }


}

// pix
// https://brown-bizarre-butterfly-586.mypinata.cloud/ipfs/bafybeif7jrpdtx6xlc3xcwfxm3zjahy4q7h3znx6avyxkg4yeumh3iftqy
// metadate
// https://brown-bizarre-butterfly-586.mypinata.cloud/ipfs/bafkreiagurvznc7aiphq47kpweuoju5eh5rq6cqcxcp66faotcym6s2pgq