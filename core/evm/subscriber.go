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
	"time"
)

type Listener struct {
	conn     *ethclient.Client
	contract common.Address
	finality uint64
}

func NewListener(client *ethclient.Client, contractAddr string, blockFinality uint) Listener {
	return Listener{
		conn:     client,
		contract: common.HexToAddress(contractAddr),
		finality: uint64(blockFinality),
	}
}

func (l Listener) Listen(ctx context.Context, ch chan event.Rx) error {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.contract},
	}

	logs := make(chan types.Log)
	sub, err := l.conn.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}
	// TODO add NFT contract subscriber

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			// TODO DIGEST IS BLOCKING THE EVENT SUBSCRIPTION | THINK IF SECURITY MARGIN SHOULD BE DONE ELSEWHERE
			action, err := l.digestEvent(ctx, vLog)
			if err != nil {
				return err
			}
			ch <- event.Rx{
				TxHash: vLog.TxHash,
				Action: action,
			}
		}
	}
}

func (l Listener) digestEvent(ctx context.Context, vLog types.Log) (string, error) {
	switch vLog.Topics[0].Hex() {
	case nftCustodySigHash.Hex():
		contractAbi, err := abi.JSON(strings.NewReader(string(contract.CustosialVaultMetaData.ABI)))
		if err != nil {
			return "", err
		}

		// TODO rework SC and add remaining events
		var nftCustody NftCustodyEvent
		err = contractAbi.UnpackIntoInterface(&nftCustody, nftCustodyEventName, vLog.Data)
		if err != nil {
			return "", err
		}
		nftCustody.TokenId = big.NewInt(vLog.Topics[1].Big().Int64())

		// TODO make it non blocking
		for {
			currentBlock, err := l.conn.BlockNumber(context.Background())
			if err != nil {
				return "", err
			}

			if currentBlock-vLog.BlockNumber >= l.finality {
				return event.MintAction, nil
			}
			time.Sleep(time.Second * 5)
		}
	default:
		return "", fmt.Errorf("unknown event signature hash: %v", vLog.Topics[0].Hex())
	}
}
