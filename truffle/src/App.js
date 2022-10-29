import React, { Fragment, useState } from "react";
import logo from './logo.svg';
import './App.css';
import logging from './bridge'
import {rejects} from "assert";
import Web3 from "web3";
import CustodialVault from "../build/contracts/ERC721CustodialVault.json"
import ERC721Token from "../build/contracts/ERC721Token.json"


// const ethereumButton = document.querySelector('.enableEthereumButton');
// const showAccount = document.querySelector('.showAccount');
//
// ethereumButton.addEventListener('click', () => {
//   getAccount();
// });
//
// async function getAccount() {
//   const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
//   const account = accounts[0];
//   showAccount.innerHTML = account;
// }
const Dropdown = ({ label, value, options, onChange }) => {
    return (
        <label>
            {label}
            <select value={value} onChange={onChange}>
                {options.map((option) => (
                    <option value={option.value}>{option.label}</option>
                ))}
            </select>
        </label>
    );
};

function App() {
    const web3 = new Web3(window.ethereum);

    const options = [
        { label: 'Ethereum', value: 'ethereum' },
        { label: 'Polygon', value: 'polygon' },
        { label: 'Binance', value: 'binance' },
    ];
    const [sourceValue, setSourceValue] = React.useState('ethereum');
    const handleSourceChange = (event) => {
        setSourceValue(event.target.value);
    };
    const [destinationValue, setDestinationValue] = React.useState('polygon');
    const handleDestinationChange = (event) => {
        setDestinationValue(event.target.value);
    };

    const [addressValue, setAddressValue] = React.useState('');
    const handleConnection = (event) => {
        setAddressValue(event.target.value);
    };

    const address = async () => {
        const accounts = web3.eth.requestAccounts().then(console.log);
        return await web3.eth.getBalance(accounts[0]).address;
    }


    const connect = async () => {
        return new Promise(async (resolve, reject)=> {
            try {
                await window.ethereum.request({method: "eth_requestAccounts"})
                resolve(web3)
                setAddressValue(await address());
            } catch (error) {
                reject(error)
            }
        })
    }

    const providerRpc = async (chainId) => {
        const goerliChainId = Number(process.env.REACT_APP_goerliChainId);
        const mumbaiChainId = Number(process.env.REACT_APP_mumbaiChainId);
        const bsctChainId = Number(process.env.REACT_APP_bsctChainId);
        console.log("Ci: ", chainId);
        console.log("G: ", goerliChainId);
        console.log("M: ", mumbaiChainId);
        console.log("B: ", bsctChainId);
        let provider;
        switch(chainId) {
            // case goerliChainId:
            //     provider = process.env.REACT_APP_goerliRpc;
            //     break;
            case mumbaiChainId:
                provider = process.env.REACT_APP_mumbaiRpc;
                break;
            case bsctChainId:
                provider = process.env.REACT_APP_bsctRpc;

                break;
            default:
            throw new Error("unsupported network")
        }
        return provider
    }

    const retrieveTokens = async () => {
        const networkId = await web3.eth.net.getId()
        const rpc = await providerRpc(networkId);
        console.log("RPC: ", rpc)
        const provider = new Web3(rpc);
        const networkData = ERC721Token.networks[networkId];
        console.log("PROVIDER ", provider)
        console.log("NETWORKDATA", networkData)

        const accounts = await web3.eth.requestAccounts();
        if(networkData) {
            console.log("network data")
            const ERC721TokenAbi = ERC721Token.abi;
            const address = networkData.address;
            const contract = new web3.eth.Contract(ERC721TokenAbi, address);
            const walletTokens = await contract.methods.walletOfOwner(accounts[0]).call();
            console.log(`${accounts[0]} tokens: ${walletTokens}`);
            // TODO print tokens in front (must be selectable)
        }

    }

    // TODO ask to switch network when mismatch between wallet network and select value

    // TODO
    const transfer = async () => {
        //https://ethereum.stackexchange.com/questions/91510/call-contract-view-method-from-web3

        // const provider = new Web3("https://goerli.infura.io/v3/72902afeffe44eeba2ed93d798f87ebd");
        // const networkId = await web3.eth.net.getId()
        // const networkData = CustodialVault.networks[networkId]
        // const accounts = await web3.eth.requestAccounts();
        // if(networkData) {
        //     const custodialVaultAbi = CustodialVault.abi;
        //     const address = networkData.address;
        //     const contract = new web3.eth.Contract(custodialVaultAbi, address);
        //     // TODO probar en interactive mode con truffle
        //     const tx = await contract.methods.reteinToken(tokenId).send({from: accounts[0], value: 1000});
        // }
    }

    // logging()
    return (
        <Fragment>
            <button className="enableEthereumButton" onClick={connect}>Connect wallet</button>
            <h2>Wallet Address: <span className="showAccount">{addressValue} </span></h2>
            <div>
                <Dropdown
                    label="Source Blockchain"
                    options={options}
                    value={sourceValue}
                    onChange={handleSourceChange}
                />
                <p>These are your ERC721 tokens on {sourceValue} </p>
                <button className="retrieveTokens" onClick={retrieveTokens}>Retrieve Tokens</button>
            </div>
            <div>
                <Dropdown
                    label="Destination Blockchain"
                    options={options}
                    value={destinationValue}
                    onChange={handleDestinationChange}
                />
                <p>The destination blockchain is {destinationValue} </p>
                <button className="transferToken" onClick={transfer}>Transfer Token</button>
            </div>

        </Fragment>
);
}

export default App;
