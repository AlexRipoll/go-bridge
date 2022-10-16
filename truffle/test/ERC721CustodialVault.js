const ERC721Token = artifacts.require("ERC721Token");
const ERC721CustodialVault = artifacts.require("ERC721CustodialVault");

contract("ERC721CustodialVault", (accounts) => {
    it("should transfer token with id 4 to vault", async () => {
        const tokenId = 4;
        const ERC721TokenInstance = await ERC721Token.new();
        const ERC721CustodialVaultInstance = await ERC721CustodialVault.new(ERC721TokenInstance.address);
        const custodialVaultAddress = ERC721CustodialVaultInstance.address;

        let vaultTokenOwner = await ERC721CustodialVaultInstance.holdCustody(4);
        assert.equal(vaultTokenOwner, 0, `expecting 0 token, got ${vaultTokenOwner}`);
        let custodialVaultTokens = (await ERC721TokenInstance.walletOfOwner(custodialVaultAddress)).length;
        assert.equal(custodialVaultTokens, 0, `expecting 0 tokens, got ${custodialVaultTokens}`);

        await ERC721TokenInstance.mint(accounts[0], tokenId);

        await ERC721TokenInstance.approve(ERC721CustodialVaultInstance.address, tokenId);
        const token = await ERC721CustodialVaultInstance.retainToken(tokenId, {from: accounts[0], value: 1000});

        const tokens = await ERC721TokenInstance.walletOfOwner(accounts[0]);
        assert.equal(tokens.length, 0, `expecting 0 token, got ${tokens.length}`);
        vaultTokenOwner = await ERC721CustodialVaultInstance.holdCustody(4);
        assert.equal(vaultTokenOwner, accounts[0], `expecting the token owner to be ${accounts[0]}, got ${vaultTokenOwner}`);
        custodialVaultTokens = (await ERC721TokenInstance.walletOfOwner(custodialVaultAddress)).length;
        assert.equal(custodialVaultTokens, 1, `expecting 1 tokens, got ${custodialVaultTokens}`);
    });
});
