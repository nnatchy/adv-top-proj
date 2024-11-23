// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract UserNFT is ERC721, Ownable {
    uint256 public nextTokenId;

    // Pass msg.sender to the Ownable constructor
    constructor() ERC721("UserNFT", "UNFT") Ownable(msg.sender) {
        nextTokenId = 0;
    }

    function mint(address to) external onlyOwner {
        _safeMint(to, nextTokenId);
        nextTokenId++;
    }

    function burn(uint256 tokenId) external {
    require(
        ownerOf(tokenId) == msg.sender || msg.sender == owner(),
        "Only the token owner or contract owner can burn the token"
    );
    _burn(tokenId);
}

}