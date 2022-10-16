var ERC721Token = artifacts.require("ERC721Token");
var ERC721CustodialVault = artifacts.require("ERC721CustodialVault");

module.exports = function(deployer) {
    // Only deploy ENS if there's not already an address already.
    // i.e., don't deploy if we're using the canonical ENS address,
    // but do deploy it if we're on a test network and ENS doesn't exist.
    deployer.deploy(ERC721Token).then(function() {
        return deployer.deploy(ERC721CustodialVault, ERC721Token.address);
    });
};
