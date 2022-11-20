# go-bridge

Simple proof of concept cross-chain brigde (EVM compatible only). It has only one relay for communicating two 
different blockchain networks.


## Run it locally

### Frontend and Contracts


Install frontend dependencies
    
    cd truffle
    npm install

Compile smart contracts
    
    truffle compile


Start truffle development environment
    
    truffle develop

Start local ganache instances (open one terminal for each)
    
    ganache --chain.chainId 97 -i 97 -b 1 -m "mimic twin dismiss topic special ranch decline hen soap buzz analyst exact"
    ganache --chain.chainId 5 -i 5 -b 1 -m "mimic twin dismiss topic special ranch decline hen soap buzz analyst exact"

Deploy smart contracts

    truffle migrate --reset --network develop
    truffle migrate --reset --network develop2
    truffle migrate --reset --network develop3

Rename .env.example to .env (change any value you want but keep in mind you will need to update truffle-config too)

### Relayer

Download relay dependencies
  
    go mod download

#### Config.yaml

Rename config.yaml.dist to config.yaml

Set privateKey with the smart contract owner private key (the one used for deploying the contracts => when running truffle develop a list of pvks is printed, copy the first one)

Set `erc721TokenAddress` and `custodialVaultAddress` for each network (contract addresses can be found in the network field in `ERC721CustodialVault.json` and `ERC721Token.json` files located in truffle/build/contracts folder)


### CLI commands

#### Minting ERC721 tokens (will mint a token in the polygon local network (localhost:9545))

    go run cmd/go-bridgecli/main.go --fn mint <wallet address to recieve the token> <token id>

#### Transfer token 

    go run cmd/go-bridgecli/main.go --fn transfer <origin wallet address> <destination wallet address> <token id>


#### Check wallet tokens
    
    go run cmd/go-bridgecli/main.go --fn wallet <wallet address>


