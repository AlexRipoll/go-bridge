package v1

import (
	"context"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunEvmEventListener(config config.Config, ch chan event.Rx) error {
	client, err := ethclient.Dial(config.Networks[config.MainNetwork].Ws)
	if err != nil {
		return err
	}

	sub := evm.NewListener(
		client,
		config.Networks[config.MainNetwork].Contracts["vault"],
		config.Networks[config.MainNetwork].BlockFinality,
	)
	if err := sub.Listen(context.Background(), ch); err != nil {
		return err
	}

	return nil
}
