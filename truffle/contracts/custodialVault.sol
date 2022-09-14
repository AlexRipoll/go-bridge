// SPDX-License-Identifier: MIT

pragma solidity 0.8.16;

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract NFTCustodialVault is IERC721Receiver, ReentrancyGuard, Ownable {

    uint256 public txCost = 0.000075 ether;

    struct CustodyVault {
        uint256 tokenId;
        address holder;
    }

    mapping(uint256 => CustodyVault) public holdCustody;

    event NFTCustody (uint256 indexed tokenId, address holder);

    ERC721Enumerable nft;

    // smart contract address to bind
    constructor(ERC721Enumerable _nft) {
        nft = _nft;
    }

    /**
    * @dev Safely transfers `tokenId` token from the owner's wallet to the contract, and retains it until it is ask
    * to be released to the owner's wallet.
    *
    * this function needs a minimum amount `txCost` of native tokens to be executed.
    *
    * Requirements:
    *
    * - `tokenId` token must exist and be owned by `from`.
    * - only the NFT owner can transfer the token to the contract vault.
    * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
    *
    * Emits a {NFTCustody} event.
    */
    function retainNFT(uint256 tokenId) public payable nonReentrant {
        require(msg.value == txCost, "Not enough balance to complete transaction.");
        require(nft.ownerOf(tokenId) == msg.sender, "you must be the NFT owner for executing this action");
        require(holdCustody[tokenId].tokenId == 0, "NFT already stored");
        holdCustody[tokenId] = CustodyVault(tokenId, msg.sender);
        nft.transferFrom(msg.sender, address(this), tokenId);
        emit NFTCustody(tokenId, msg.sender);
    }

    function updateOwner(uint256 tokenId, address newHolder) public nonReentrant onlyOwner() {
        holdCustody[tokenId] = CustodyVault(tokenId, newHolder);
        emit NFTCustody(tokenId, newHolder);
    }

    /**
    * @dev Releases a `tokenId` token hold in the `holdCustody` vault and transfers it to its owner's `wallet`.
    * to be released to the owner's wallet.
    *
    * Requirements:
    *
    * - only can be called by the contract owner.
    *
    */
    function releaseNFT(uint256 tokenId, address wallet) public nonReentrant onlyOwner() {
        nft.transferFrom(address(this), wallet, tokenId);
        delete holdCustody[tokenId];
    }

    function emergencyDelete(uint256 tokenId) public nonReentrant onlyOwner() {
        delete holdCustody[tokenId];
    }

    function onERC721Received(
        address,
        address from,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        require(from == address(0x0), "Cannot Receive NFTs Directly");
        return IERC721Receiver.onERC721Received.selector;
    }

    function withdraw() public payable onlyOwner() {
        require(payable(msg.sender).send(address(this).balance));
    }

}
