// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC721CustodialVault is IERC721Receiver, ReentrancyGuard, Ownable {

    uint256 public txCost = 1000 wei;

    mapping(uint256 => address) public holdCustody;

    event TokenCustody (uint256 indexed tokenId, address holder, uint256 destinationChainId);

    ERC721Enumerable erc721token;

    // smart contract address to bind
    constructor(ERC721Enumerable _erc721token) {
        erc721token = _erc721token;
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
    * - only the Token owner can transfer the token to the contract vault.
    * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
    *
    * Emits a {TokenCustody} event.
    */
    function retainToken(uint256 tokenId, uint256 destination) public payable nonReentrant {
        require(msg.value >= txCost, "Not enough balance to complete transaction.");
        require(erc721token.ownerOf(tokenId) == msg.sender, "you must be the Token owner for executing this action");
        require(holdCustody[tokenId] == address(0), "Token already stored");
        holdCustody[tokenId] = msg.sender;
        erc721token.transferFrom(erc721token.ownerOf(tokenId), address(this), tokenId);
        emit TokenCustody(tokenId, msg.sender, destination);
    }

    function updateOwner(uint256 tokenId, address newHolder) public nonReentrant onlyOwner() {
        holdCustody[tokenId] = newHolder;
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
    function releaseToken(uint256 tokenId, address wallet) public nonReentrant onlyOwner() {
        erc721token.transferFrom(address(this), wallet, tokenId);
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
        require(from == address(0x0), "Cannot Receive Tokens Directly");
        return IERC721Receiver.onERC721Received.selector;
    }

    function withdraw() public payable onlyOwner() {
        require(payable(msg.sender).send(address(this).balance));
    }
}
