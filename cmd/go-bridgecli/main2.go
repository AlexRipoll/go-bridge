package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core"
	evm2 "github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"strconv"
)

func main() {
	fn := flag.Bool("fn", false, "executes a function")
	deploy := flag.Bool("deploy", false, "deploy contract abis")

	flag.Parse()

	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var custodian evm2.Custodian
	bridgers := make(map[string]evm2.Erc721Token)
	for name, network := range config.Networks {
		conn, err := ethclient.Dial(network.Http)
		if err != nil {
			log.Fatal(err)
		}
		for contract, address := range network.Contracts {
			switch contract {
			case "vault":
				custodian, err = evm2.NewCustodian(conn, address, config, name)
				if err != nil {
					log.Fatal(err)
				}
			case "bridge":
				bridger, err := evm2.NewBridger(conn, address, config, name)
				if err != nil {
					log.Fatal(err)
				}
				bridgers[name] = bridger
			default:
				log.Fatal("unknown network type")
			}
		}

	}
	bridge, err := core.NewBridge(config.MainNetwork, custodian, bridgers)
	if err != nil {
		log.Fatal(err)
	}

	if *fn {
		switch flag.Arg(0) {
		case "mint":
			walletAddr := flag.Arg(1)

			networkName := config.MainNetwork
			network := config.Networks[networkName]
			conn, err := ethclient.Dial(network.Http)
			if err != nil {
				log.Fatal(err)
			}
			minter, err := evm2.NewBridger(conn, network.Contracts["bridge"], config, networkName)
			if err != nil {
				log.Fatal(err)
			}
			bridgers[networkName] = minter

			tx, err := minter.Mint(context.Background(), walletAddr, big.NewInt(6))
			if err != nil {
				log.Fatal(" error minting token: ", err)
			}
			fmt.Println(fmt.Sprintf("%#v", tx))
		case "transfer":
			destinationBlockchain := flag.Arg(1)
			originBlockchain := flag.Arg(2)
			walletAddr := flag.Arg(3)
			tokenIdArg := flag.Arg(4)
			tokenId, err := strconv.Atoi(tokenIdArg)
			if err != nil {
				log.Fatal(err)
			}

			err = bridge.TransferNFT(context.Background(), destinationBlockchain, originBlockchain, walletAddr, big.NewInt(int64(tokenId)))
			if err != nil {
				log.Fatal(" error transferring token: ", err)
			}
			fmt.Println(">> token transfer success")
		case "wallet":
			walletAddr := flag.Arg(0)
			blockchain := flag.Arg(1)

			tokens, err := bridge.WalletTokens(walletAddr, blockchain)
			if err != nil {
				log.Fatalf("error fetching tokens: %v", err)
			}
			fmt.Println("Token Ids: ", tokens)
		default:
			log.Fatalf("unknow method: %s", flag.Arg(0))
		}
	}

	if *deploy {
		if err = bridge.Deploy(context.Background()); err != nil {
			log.Fatal(errors.Wrap(err, "error deploying contract"))
		}
		return
	}
}
