// SPDX-License-Identifier: MIT

pragma solidity 0.8.16;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";

contract NFT is ERC721Enumerable, Ownable {

    using Strings for uint256;
    string public baseURI;
    string public baseExtension = ".json";
    bool public paused = false;

    event NFTMint (uint256 indexed tokenId, address holder);
    event NFTBurn (uint256 indexed tokenId, address holder);

    constructor() ERC721("NFT Collection", "NCN") {
    }

    function _baseURI() internal view virtual override returns (string memory) {
    return "ipfs://QmYB5uWZqfunBq7yWnamTqoXWBAHiQoirNLmuxMzDThHhi/";

    }
    
    function bridgeMint(address to, uint256 tokenId) public virtual onlyOwner() {
        _mint(to, tokenId);
        emit NFTMint(tokenId, msg.sender);
    }

    function bridgeBurn(uint256 tokenId) public virtual onlyOwner() {
        _burn(tokenId);
        emit NFTBurn(tokenId, msg.sender);
    }

    function walletOfOwner(address _owner)
        public
        view
        returns (uint256[] memory)
        {
            uint256 ownerTokenCount = balanceOf(_owner);
            uint256[] memory tokenIds = new uint256[](ownerTokenCount);
            for (uint256 i; i < ownerTokenCount; i++) {
                tokenIds[i] = tokenOfOwnerByIndex(_owner, i);
            }
            return tokenIds;
        }
    
        
        function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory) {
            require(
                _exists(tokenId),
                "ERC721Metadata: URI query for nonexistent token"
                );
                
                string memory currentBaseURI = _baseURI();
                return
                bytes(currentBaseURI).length > 0 
                ? string(abi.encodePacked(currentBaseURI, tokenId.toString(), baseExtension))
                : "";
        }
        // only owner

        function setBaseURI(string memory _newBaseURI) public onlyOwner() {
            baseURI = _newBaseURI;
        }
        
        function setBaseExtension(string memory _newBaseExtension) public onlyOwner() {
            baseExtension = _newBaseExtension;
        }
        
        function pause(bool _state) public onlyOwner() {
            paused = _state;
        }

}