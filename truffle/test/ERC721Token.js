const ERC721Token = artifacts.require("ERC721Token");

contract("ERC721Token", (accounts) => {
    it("should mint token with id 4 and being added to accounts[0]", async () => {
        const ERC721TokenInstance = await ERC721Token.deployed();

        await ERC721TokenInstance.mint(accounts[0], 4);
        const tokens = await ERC721TokenInstance.walletOfOwner(accounts[0]);
        assert.equal(tokens.length, 1, `expecting 1 token, got ${tokens.length}`);
        assert.equal(tokens[0].toString(), 4, `expecting token id 4, got ${tokens[0].toString()}`);
    });
    it("should burn token with id 5 and being removed from accounts[0]", async () => {
        const ERC721TokenInstance = await ERC721Token.deployed();

        await ERC721TokenInstance.mint(accounts[0], 5);
        let tokens = await ERC721TokenInstance.walletOfOwner(accounts[0]);
        assert.equal(tokens.length, 2, `expecting 2 token, got ${tokens.length}`);

        await ERC721TokenInstance.burn(5);
        tokens = await ERC721TokenInstance.walletOfOwner(accounts[0]);
        assert.equal(tokens.length, 1, `expecting 1 token, got ${tokens.length}`);
        assert.equal(tokens[0].toString(), 4, `expecting token id 4, got ${tokens[0].toString()}`);
    });
});
