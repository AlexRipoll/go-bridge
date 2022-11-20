require("dotenv").config({ path: '.env' });

const HDWalletProvider = require('@truffle/hdwallet-provider');
// const mnemonic = process.env["mnemonic"];
const privateKey = process.env["privateKey"];

module.exports = {
    networks: {
        develop: {
            host: "127.0.0.1",
            port: 9545,
            network_id: "1337",
            blockTime: 1,
            websockets: true
        },
        develop2: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "97",
            blockTime: 1,
            websockets: true
        },
        goerli: {
            provider: function() {
                return new HDWalletProvider(privateKey, "https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161");
            },
            network_id: '5',
        },
        mumbai: {
            provider: function() {
                return new HDWalletProvider(privateKey, "https://matic-mumbai.chainstacklabs.com");
            },
            network_id: '80001',
        },
        bsct: {
            provider: function() {
                return new HDWalletProvider(privateKey, "https://data-seed-prebsc-1-s1.binance.org:8545/");
            },
            network_id: '97',
        },
    },
    compilers: {
        solc: {
            version: "^0.8.16"
        }
    }
};
