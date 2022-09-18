package main

import (
	"context"
	"flag"
	"github.com/AlexRipoll/go-bridge/blockchain/config"
	"github.com/AlexRipoll/go-bridge/blockchain/core"
	"github.com/AlexRipoll/go-bridge/blockchain/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"log"
)

func main() {
	deploy := flag.Bool("deploy", false, "deploy contract abis")

	flag.Parse()

	if *deploy {
		var custodian evm.Custodian
		bridgers := make(map[string]evm.Bridger)
		for _, n := range config.Networks {
			conn, err := ethclient.Dial(n.Url)
			if err != nil {
				log.Fatal(err)
			}

			switch n.Type {
			case "custodian":
				custodian, err = evm.NewCustodian(conn, n.Contracts["vault"], config, n.Name)
				if err != nil {
					log.Fatal(err)
				}
			case "bridger":
				bridger, err := evm.NewBridger(conn, n.Contracts["bridge"], config, n.Name)
				if err != nil {
					log.Fatal(err)
				}
				bridgers[n.Name] = bridger
			default:
				log.Fatal("unknown network type")
			}

		}
		bridge, err := core.NewBridge(custodian, bridgers)
		if err != nil {
			log.Fatal(err)
		}
		if err = bridge.Deploy(context.Background()); err != nil {
			log.Fatal(errors.Wrap(err, "error deploying contract"))
		}
	}
}