package main

import (
	"context"
	"fmt"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/AlexRipoll/go-bridge/sys/storage"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("initializing process...")

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewLevelDB("./leveldb")
	if err != nil {
		log.Fatal(err)
	}

	//goerli := config.Goerli
	mumbai := config.Mumbai
	bsct := config.Bsct

	//goerliClient := NewClient(
	//	goerli.Http,
	//	goerli.ChainId,
	//	goerli.BlockFinality,
	//	config.PrivateKey,
	//	goerli.Contracts.CustodialVaultAddress,
	//	goerli.Contracts.ERC721TokenAddress,
	//	)

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

	clients := make(map[uint64]evm.Client)
	clients[mumbai.ChainId] = mumbaiClient
	clients[bsct.ChainId] = bsctClient

	releaser := evm.Releaser{
		Clients: clients,
	}

	//goerliWs, err := ethclient.Dial(goerli.Ws)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Infof("connection established to %v", goerli.Ws)

	//ethListner := evm.NewListener(goerliWs, goerli.Contracts.CustodialVaultAddress, releaser)
	//go ethListner.Listen(context.Background())

	go evm.RecoverFromShutDown(context.Background(), db, releaser)

	fmt.Println(mumbai.Ws)
	polWs, err := ethclient.Dial(mumbai.Ws)
	if err != nil {
		log.Fatal("----", err)
	}
	log.Infof("connection established to %v", mumbai.Ws)

	polygonListner := evm.NewListener(polWs, mumbai.Contracts.CustodialVaultAddress, mumbai.ChainId, releaser, db)
	go polygonListner.Listen(context.Background())

	bscWs, err := ethclient.Dial(bsct.Ws)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("connection established to %v", bsct.Ws)

	bscListner := evm.NewListener(bscWs, bsct.Contracts.CustodialVaultAddress, bsct.ChainId, releaser, db)
	go bscListner.Listen(context.Background())

	ch := make(chan struct{})

	<-ch
}

func NewClient(http string, chainId, finality uint64, privateKey, custodialVaultAddress, erc721TokenAddress string) evm.Client {
	conn, err := ethclient.Dial(http)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("connection established to %v", http)

	transactor := evm.NewTransactor(conn, privateKey)
	custodian, err := evm.NewErc721CustodialVaultContract(custodialVaultAddress, transactor, conn)
	if err != nil {
		log.Fatalf("NewErc721CustodialVaultContract error: %v", err)
	}
	erc721Token, err := evm.NewErc721TokenContract(erc721TokenAddress, transactor, conn)
	if err != nil {
		log.Fatalf("NewErc721TokenContract error: %v", err)
	}
	return evm.NewClient(conn, chainId, finality, custodian, erc721Token, transactor)
}
