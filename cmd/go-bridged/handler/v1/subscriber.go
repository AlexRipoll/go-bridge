package v1

import (
	"context"
	"github.com/AlexRipoll/go-bridge/blockchain/config"
	"github.com/AlexRipoll/go-bridge/blockchain/core/evm"
	"github.com/AlexRipoll/go-bridge/blockchain/core/scanner"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunEventListener(config config.Config, ch chan scanner.EventRx) error {
	client, err := ethclient.Dial(config.Networks[config.MainNetwork].Ws)
	if err != nil {
		return err
	}
	sub := evm.NewSubscriber(client, config.Networks[config.MainNetwork].Contracts["vault"])
	if err := sub.ListenEvents(context.Background(), ch); err != nil {
		return err
	}

	return nil
}