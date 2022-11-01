package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

func main() {
	fn := flag.Bool("fn", false, "executes a function")

	flag.Parse()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}


	mumbai := config.Mumbai
	bsct := config.Bsct

	log.Info("Chain Id: ", mumbai.ChainId)
	log.Info("RPC: ", mumbai.Http)
	log.Info("ERC721TokenAddress: ", mumbai.Contracts.ERC721TokenAddress)

	log.Info("Chain Id: ", bsct.ChainId)
	log.Info("RPC: ", bsct.Http)
	log.Info("ERC721TokenAddress: ", bsct.Contracts.ERC721TokenAddress)

	conn, err := ethclient.Dial(mumbai.Http)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("connection established to %v", mumbai.Http)
	transactor := evm.NewTransactor(conn, config.PrivateKey)
	erc721Token, err := evm.NewErc721TokenContract(mumbai.Contracts.ERC721TokenAddress, transactor, conn)
	if err != nil {
		log.Fatalf("NewErc721TokenContract error: %v", err)
	}

	//conn, err := ethclient.Dial(bsct.Http)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Infof("connection established to %v", mumbai.Http)
	//transactor := evm.NewTransactor(conn, config.PrivateKey)
	//erc721Token, err := evm.NewErc721TokenContract(bsct.Contracts.ERC721TokenAddress, transactor, conn)
	//if err != nil {
	//	log.Fatalf("NewErc721TokenContract error: %v", err)
	//}

	service := core.NewService(erc721Token)

	if *fn {
		switch flag.Arg(0) {
		case "mint":
			walletAddress := flag.Arg(1)
			tokenIdArg := flag.Arg(2)
			tokenId, err := strconv.Atoi(tokenIdArg)
			if err != nil {
				log.Fatalf("%v is not an integer", tokenIdArg)
			}

			tx, err := service.MintERC721Token(context.Background(), walletAddress, big.NewInt(int64(tokenId)))
			if err != nil {
				log.Fatal(" error minting token: ", err)
			}
			fmt.Println(fmt.Sprintf("%#v", tx))
		case "wallet":
			walletAddress := flag.Arg(1)

			tokens, err := service.WalletTokens(context.Background(), walletAddress)
			if err != nil {
				log.Fatalf("error fetching tokens: %v", err)
			}
			fmt.Println("Token Ids: ", tokens)
		default:
			log.Fatalf("unknow method: %s", flag.Arg(0))
		}
	}

}
