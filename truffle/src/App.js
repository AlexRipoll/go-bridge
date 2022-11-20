import React, { Fragment, useState } from "react";
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import Modal from 'react-bootstrap/Modal';
import Form from 'react-bootstrap/Form';
import 'bootstrap/dist/css/bootstrap.min.css';
import logo from './logo.svg';
import './App.css';
import logging from './bridge'
import {rejects} from "assert";
import Web3 from "web3";
import CustodialVault from "../build/contracts/ERC721CustodialVault.json"
import ERC721Token from "../build/contracts/ERC721Token.json"



const ethereumNetId = process.env.REACT_APP_ethereumNetId;
const polygonNetId = process.env.REACT_APP_polygonNetId;
const binanceNetId = process.env.REACT_APP_binanceNetId;

const Dropdown = ({ label, value, options, onChange }) => {
    return (
        <Form.Group>
            <Form.Label>{label}</Form.Label>
            <Form.Select value={value} onChange={onChange}>
                {options.map((option) => (
                    <option value={option.value}>{option.label}</option>
                ))}
            </Form.Select>
        </Form.Group>
    );
};

const TokenCard = function ({tokenId, selectedId, onClick}) {
    const isSelected = selectedId === tokenId;
    return (
        <Card style={{ width: '8rem' }}>
            <Card.Img variant="top" src="https://via.placeholder.com/150" />
            <Card.Body>
                <Card.Title>{tokenId}</Card.Title>
                <Button disabled={isSelected} onClick={onClick}>Select</Button>
            </Card.Body>
        </Card>
    );
}

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

    const [retrievedTokens, setRetrievedTokens ] = useState([]);

    const [selectedToken, setSelectedToken] = useState();

    const address = async () => {
        const accounts = web3.eth.requestAccounts().then(console.log);
        return await web3.eth.getBalance(accounts[0]).address;
    }

    const isConnected = !!addressValue;

    console.log({selectedToken, addressValue, sourceValue, destinationValue})

    const connect = async () => {
        return new Promise(async (resolve, reject)=> {
            try {
                const [address] = await window.ethereum.request({method: "eth_requestAccounts"})
                setAddressValue(address);
                //resolve(web3)

            } catch (error) {
                reject(error)
            }
        })
    }

    const providerRpc = async (chainId) => {
        let provider;
        switch(chainId) {
            case ethereumNetId:
                provider = process.env.REACT_APP_ethereumRpc;
                break;
            case binanceNetId:
                provider = process.env.REACT_APP_binanceRpc;
                break;
            case polygonNetId:
                provider = process.env.REACT_APP_polygonRpc;
                break;
            default:
            throw new Error("unsupported network")
        }
        return provider
    }

    const checkNetwork = async function(selectedNetId) {
        // Check if MetaMask is installed
        // MetaMask injects the global API into window.ethereum
        if (window.ethereum) {
            try {
                // check if the chain to connect to is installed
                await window.ethereum.request({
                    method: 'wallet_switchEthereumChain',
                    params: [{ chainId: `0x${selectedNetId.toString(16)}` }], // chainId must be in hexadecimal numbers
                });
            } catch (error) {
                // This error code indicates that the chain has not been added to MetaMask
                // if it is not, then install it into the user MetaMask
                if (error.code === 4902) {
                    try {
                        await window.ethereum.request({
                            method: 'wallet_addEthereumChain',
                            params: [
                                {
                                    chainId: '0x61',
                                    rpcUrl: 'https://data-seed-prebsc-1-s1.binance.org:8545/',
                                },
                            ],
                        });
                    } catch (addError) {
                        console.error(addError);
                    }
                }
                console.error(error);
            }
        } else {
            // if no window.ethereum then MetaMask is not installed
            alert('MetaMask is not installed. Please consider installing it: https://metamask.io/download.html');
        }
    }

    const retrieveTokens = async () => {
        let networkId = await getSelectedNetworkId(sourceValue);

        await checkNetwork(networkId);
        const rpc = await providerRpc(networkId);
        const provider = new Web3(rpc);
        const networkData = ERC721Token.networks[networkId];

        if(networkData) {
            const ERC721TokenAbi = ERC721Token.abi;
            const address = networkData.address;
            const contract = new web3.eth.Contract(ERC721TokenAbi, address);
            const walletTokens = await contract.methods.walletOfOwner(addressValue).call();
            console.log(`>>>> ${addressValue} tokens: ${walletTokens}`);

            const vaultNetworkData = CustodialVault.networks[networkId];
            const vaultTokens = await contract.methods.walletOfOwner(vaultNetworkData.address).call();
            console.log(`>>>> custodial Vault tokens: ${vaultTokens}`);
            // TODO print tokens in front (must be selectable)
            setRetrievedTokens(walletTokens);
            console.log(`>>>> wallet tokens: ${walletTokens}`);
        }

    }

    const printTokens = async (list) => {
        const cards = list.map((item) => {
            <li>newCard(item.id)</li>
        });
        return cards;
    }


    const getSelectedNetworkId = async (selectedValue) => {
        let networkId;
        if (selectedValue == "polygon") {
            networkId = polygonNetId;
        } else if (selectedValue == "binance") {
            networkId = binanceNetId;
        } else if (selectedValue == "ethereum") {
            networkId = binanceNetId;
        }
        console.log(">>>> net id: ", networkId);

        return networkId;
    }

    const transfer = async () => {
        //https://ethereum.stackexchange.com/questions/91510/call-contract-view-method-from-web3
        // https://web3js.readthedocs.io/en/v1.2.11/web3-eth-contract.html

        const networkId = await getSelectedNetworkId(sourceValue);

        await checkNetwork(networkId);
        const rpc = await providerRpc(networkId);
        const provider = new Web3(rpc);
        const networkData = CustodialVault.networks[networkId];
        const erc721TokenNetworkData = ERC721Token.networks[networkId];

        const destination = await getSelectedNetworkId(destinationValue);

        const accounts = await web3.eth.requestAccounts();
        if (networkData) {
            const CustodialVaultAbi = CustodialVault.abi;
            const address = networkData.address;
            const contract = new web3.eth.Contract(CustodialVaultAbi, address);

            const ERC721TokenAbi = ERC721Token.abi;
            const erc721TokenAddress = erc721TokenNetworkData.address;
            const erc721TokenContract = new web3.eth.Contract(ERC721TokenAbi, erc721TokenAddress);

            const isApproved = await erc721TokenContract.methods.isApprovedForAll(accounts[0], address).call({from: accounts[0]});
            console.log(">>>> is approved: ", isApproved);
            if (!isApproved) {
                const approve = await erc721TokenContract.methods.setApprovalForAll(address, true).send({from: accounts[0], gasPrice: 30000000000});
            }

            const estimateGasPrice = await contract.methods.retainToken(selectedToken, destination).estimateGas({from: accounts[0], value: 1000});
            contract.methods.retainToken(selectedToken, destination).send({
                from: accounts[0],
                value: 1000,
                gasPrice: 30000000000,
            })
            .on('receipt', function(receipt){
                console.log(receipt)
            })
            .on('error', function(error){
                console.log("RETAIN TOKEN ERROR: ", error)
            })

            ;
        }
    }

    // logging()
    return (
        <Modal show centered >
            <Modal.Body>

            <h5>Wallet Address:</h5>
            <p className="showAccount">{ isConnected ? addressValue : 'Not Connected'} </p>

                {!!retrievedTokens?.length && <>
                    <h5>Wallet Tokens:</h5>
                    <div className={"d-flex justify-content-between gap-2 mb-4"}>
                    {retrievedTokens.map((id) => <TokenCard tokenId={id} selectedId={selectedToken} onClick={() => setSelectedToken(id)} />)}
                    </div>
                </>}


            <div className={"d-flex justify-content-between"}>
                <Dropdown
                    label="Source Blockchain"
                    options={options}
                    value={sourceValue}
                    onChange={handleSourceChange}
                    disabled={!isConnected}
                />
                <Dropdown
                    label="Destination Blockchain"
                    options={options}
                    value={destinationValue}
                    onChange={handleDestinationChange}
                    disabled={!isConnected}
                />
                <p></p>
            </div>
            </Modal.Body>
            <Modal.Footer className={"d-flex justify-content-between"}>
                {!isConnected && <Button variant="dark" className="enableEthereumButton" onClick={connect}>Connect wallet</Button>}
                { isConnected &&
                    <>
                        <Button variant="primary" className="retrieveTokens" onClick={retrieveTokens}>Retrieve Tokens</Button>
                        <Button variant="primary" className="transferToken" onClick={transfer}>Transfer Token</Button>
                    </>
                }
            </Modal.Footer>
        </Modal>
);
}

export default App;
