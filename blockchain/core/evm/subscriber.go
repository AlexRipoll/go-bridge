package evm

import (
	"context"
	"fmt"
	"github.com/AlexRipoll/go-bridge/blockchain/contract"
	"github.com/AlexRipoll/go-bridge/blockchain/core/scanner"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

type Subscriber struct {
	conn *ethclient.Client
	contract common.Address
}

func NewSubscriber(client *ethclient.Client, contractAddr string) Subscriber {

	return Subscriber{
		conn:    client,
		contract: common.HexToAddress(contractAddr),
	}
}


func (s Subscriber) ListenEvents(ctx context.Context, ch chan scanner.EventRx) error  {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{s.contract},
	}

	logs := make(chan types.Log)
	sub, err := s.conn.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			// TODO DIGEST IS BLOCKING THE EVENT SUBSCRIPTION | THINK IF SECURITY MARGIN SHOULD BE DONE ELSEWHERE
			action, err := s.digestEvent(ctx, vLog)
			if err != nil {
				return err
			}
			ch <- scanner.EventRx{
				TxHash: vLog.TxHash,
				Action: action,
			}
		}
	}
}


func (s Subscriber) digestEvent(ctx context.Context, vLog types.Log) (string, error) {
	switch vLog.Topics[0].Hex() {
	case nftCustodySigHash.Hex():
		contractAbi, err := abi.JSON(strings.NewReader(string(contract.CustosialVaultMetaData.ABI)))
		if err != nil {
			return "", err
		}

		var nftCustody NftCustodyEvent
		err = contractAbi.UnpackIntoInterface(&nftCustody, "NFTCustody", vLog.Data)
		if err != nil {
			return "", err
		}
		nftCustody.TokenId = big.NewInt(vLog.Topics[1].Big().Int64())

		for {
			currentBlock, err := s.conn.BlockNumber(context.Background())
			if err != nil {
				return "", err
			}

			// TODO mint/burn/release token
			if currentBlock - vLog.BlockNumber >= 50 {
				return "mint", nil
			}
			time.Sleep(time.Second * 5)
		}
	default:
		return "", fmt.Errorf("unknown event signature hash: %v", vLog.Topics[0].Hex())
	}
}
