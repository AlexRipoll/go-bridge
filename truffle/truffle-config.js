require("dotenv").config({ path: '../.env' });

const HDWalletProvider = require('@truffle/hdwallet-provider');
// const mnemonic = process.env["mnemonic"];
const privateKey = process.env["privateKey"];

module.exports = {
    networks: {
        // dev: {
        //     host: "127.0.0.1",
        //     port: 9545,
        //     network_id: "*",
        //     websockets: true
        // },
        goerli: {
            provider: function() {
                return new HDWalletProvider(privateKey, "https://goerli.infura.io/v3/72902afeffe44eeba2ed93d798f87ebd");
            },
            network_id: '5',
            gas: 4465030,
            gasPrice: 63363054162,
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
