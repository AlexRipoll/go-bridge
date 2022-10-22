package evm

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

var (
	ethereumChainId = *big.NewInt(5)
	polygonChainId  = *big.NewInt(80001)
	bscChainId      = *big.NewInt(59)
)

type Releaser struct {
	EthereumClient Client
	PolygonClient  Client
	BinanceClient  Client
}

func (r Releaser) releaseToken(ctx context.Context, rx event.Rx) error {
		switch rx.Destination.Int64() {
		case ethereumChainId.Int64():
			if err := r.waitForFinality(ctx, r.EthereumClient, rx); err != nil {
				log.Error(fmt.Sprintf("wait for finality error (ethereum): %v", err.Error()))
				return err
			}
			return nil
		case polygonChainId.Int64():
			if err := r.waitForFinality(ctx, r.PolygonClient, rx); err != nil {
				log.Error(fmt.Sprintf("wait for finality error (polygon): %v", err.Error()))
				return err
			}
			return nil
		case bscChainId.Int64():
			if err := r.waitForFinality(ctx, r.BinanceClient, rx); err != nil {
				log.Error(fmt.Sprintf("wait for finality error (binance): %v", err.Error()))
				return err
			}
			return nil
		default:
			return fmt.Errorf("unknown destination blockchain: %v", rx.Destination.Int64())
		}
}

func (r Releaser) waitForFinality(ctx context.Context, client Client, rx event.Rx) error {
	exists, err := client.contracts.erc721Token.Exists(ctx, rx.TokenId)
	if err != nil {
		return err
	}

	for  {
		currentBlock, err := client.Transactor.CurrentBlock(ctx)
		if err != nil {
			return err
		}

		if (currentBlock - rx.TxBlock) >= client.Finality() {
			if !exists {
				if _, err := client.contracts.erc721Token.Mint(ctx, rx.Holder, rx.TokenId); err != nil {
					return err
				}
				return nil
			}
			// check if owner is custodial vault
			if err := checkRetention(ctx, client, rx.TokenId); err != nil {
				return err
			}

			_, err = client.contracts.custodianVault.ReleaseToken(ctx, rx.Holder, rx.TokenId)
			if err != nil {
				return err
			}

			return nil
		}
		time.Sleep(time.Second * 5)
	}
}

func checkRetention(ctx context.Context, client Client, tokenId *big.Int) error {
	owner, err := client.contracts.erc721Token.OwnerOf(ctx, tokenId)
	if err != nil {
		return err
	}
	address := client.contracts.custodianVault.Address()
	if owner != address {
		return errors.New("token not retained in custody vault")
	}

	return nil
}