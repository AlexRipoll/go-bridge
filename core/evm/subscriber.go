package evm

import (
	"context"
	"fmt"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type Listener struct {
	conn     *ethclient.Client
	contract common.Address
	releaser Releaser
}

func NewListener(conn *ethclient.Client, contractAddress string, releaser Releaser) Listener {
	return Listener{
		conn:     conn,
		contract: common.HexToAddress(contractAddress),
		releaser: releaser,
	}
}

func (l Listener) Listen(ctx context.Context) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.contract},
	}

	logs := make(chan types.Log)
	sub, err := l.conn.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			tokenEvent, err := l.digestEvent(ctx, vLog)
			if err != nil {
				return err
			}

			rx := event.Rx{
				TxHash:      vLog.TxHash,
				TxBlock:     vLog.BlockNumber,
				TokenId:     tokenEvent.TokenId,
				Holder:      tokenEvent.Holder.String(),
				Destination: tokenEvent.Destination,
			}
			go l.releaser.releaseToken(ctx, rx)
		}
	}
}

func (l Listener) digestEvent(ctx context.Context, vLog types.Log) (*TokenCustodyEvent, error) {
	switch vLog.Topics[0].Hex() {
	case tokenCustodySigHash.Hex():
		contractAbi, err := abi.JSON(strings.NewReader(string(contract.CustosialVaultMetaData.ABI)))
		if err != nil {
			return nil, err
		}

		var tokenCustodyEvent TokenCustodyEvent
		err = contractAbi.UnpackIntoInterface(&tokenCustodyEvent, tokenCustodyEventName, vLog.Data)
		if err != nil {
			return nil, err
		}
		tokenCustodyEvent.TokenId = big.NewInt(vLog.Topics[1].Big().Int64())

		return &tokenCustodyEvent, nil
	default:
		return nil, fmt.Errorf("unknown event signature hash: %v", vLog.Topics[0].Hex())
	}
}