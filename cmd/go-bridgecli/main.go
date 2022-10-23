package main

import (
	"context"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	// add tests
	// update config
	// move code to service
	// recovery system (DB...)
	// logging
	// update function
	// docs

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	goerli := config.Goerli
	mumbai := config.Mumbai
	bsct := config.Bsct

	goerliClient := NewClient(
		goerli.Http,
		goerli.ChainId,
		goerli.BlockFinality,
		config.PrivateKey,
		goerli.Contracts.CustodialVaultAddress,
		goerli.Contracts.ERC721TokenAddress,
		)

	mumbaiClient := NewClient(
		mumbai.Http,
		mumbai.ChainId,
		mumbai.BlockFinality,
		config.PrivateKey,
		mumbai.Contracts.CustodialVaultAddress,
		mumbai.Contracts.ERC721TokenAddress,
	)

	bsctClient := NewClient(
		bsct.Http,
		bsct.ChainId,
		bsct.BlockFinality,
		config.PrivateKey,
		bsct.Contracts.CustodialVaultAddress,
		bsct.Contracts.ERC721TokenAddress,
	)

	releaser := evm.Releaser{
		EthereumClient: goerliClient,
		PolygonClient:  mumbaiClient,
		BinanceClient:  bsctClient,
	}

	goerliWs, err := ethclient.Dial(goerli.Ws)
	if err != nil {
		log.Fatal(err)
	}
	ethListner := evm.NewListener(goerliWs, goerli.Contracts.CustodialVaultAddress, releaser)
	ethListner.Listen(context.Background())

	polWs, err := ethclient.Dial(mumbai.Ws)
	if err != nil {
		log.Fatal(err)
	}
	polygonListner := evm.NewListener(polWs, mumbai.Contracts.CustodialVaultAddress, releaser)
	polygonListner.Listen(context.Background())

	bscWs, err := ethclient.Dial(bsct.Ws)
	if err != nil {
		log.Fatal(err)
	}
	bscListner := evm.NewListener(bscWs, bsct.Contracts.CustodialVaultAddress, releaser)
	bscListner.Listen(context.Background())
}

func NewClient(http string, chainId, finality uint64, privateKey, custodialVaultAddress, erc721TokenAddress string) evm.Client {
	conn, err := ethclient.Dial(http)
	if err != nil {
		log.Fatal(err)
	}

	transactor := evm.NewTransactor(conn, privateKey)
	custodian, err := evm.NewErc721CustodialVaultContract(custodialVaultAddress, transactor, conn)
	if err != nil {
		log.Fatalf("NewErc721CustodialVaultContract error: %v", err)
	}
	erc721Token, err := evm.NewErc721TokenContract(erc721TokenAddress, transactor, conn)
	if err != nil {
		log.Fatalf("NewErc721TokenContract error: %v", err)
	}
	return evm.NewClient(conn, chainId, finality, custodian, erc721Token)
}