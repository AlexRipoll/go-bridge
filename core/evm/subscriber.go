package evm

import (
	"context"
	"fmt"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/AlexRipoll/go-bridge/sys/storage"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

type Listener struct {
	conn     *ethclient.Client
	contract common.Address
	chainId  *big.Int
	releaser Releaser
	db storage.Db
}

func NewListener(conn *ethclient.Client, contractAddress string, chainId uint64, releaser Releaser, db storage.Db) Listener {
	return Listener{
		conn:     conn,
		contract: common.HexToAddress(contractAddress),
		chainId:  big.NewInt(int64(chainId)),
		releaser: releaser,
		db: db,
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
	log.Infof("listening events at address %v", l.contract.String())

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			log.Infof("event received...")
			tokenEvent, err := l.digestEvent(ctx, vLog)
			if err != nil {
				log.Fatalf("error digesting: %v", err)
				return err
			}

			rx := event.Rx{
				TxHash:      vLog.TxHash,
				TxBlock:     vLog.BlockNumber,
				TokenId:     tokenEvent.TokenId,
				Holder:      tokenEvent.Holder.String(),
				Origin:      l.chainId,
				Destination: tokenEvent.DestinationChainId,
			}

			buf, err := rx.Bytes()
			if err != nil {
				return err
			}

			key := make([]byte, len(l.chainId.Bytes())+len(rx.TxHash))
			key = append(l.chainId.Bytes())
			key = append(key, rx.TxHash[:]...)
			if err := l.db.Put(key, buf); err != nil {
				return err
			}
			if err := l.db.Put([]byte(LatestBlockNumber),big.NewInt(int64(rx.TxBlock)).Bytes()); err != nil {
				return err
			}

			log.Infof("event receipt %#v", l.contract.String())
			go l.releaser.releaseToken(ctx, rx)
		}
	}
}

func (l Listener) digestEvent(ctx context.Context, vLog types.Log) (*TokenCustodyEvent, error) {
	log.Infof("digesting event %v", vLog.TxHash)
	log.Infof("Topic hex is %v", vLog.Topics[0].Hex())
	switch vLog.Topics[0].Hex() {
	case TokenCustodySigHash():
		log.Infof("%v case...", TokenCustodySigHash())
		contractAbi, err := abi.JSON(strings.NewReader(string(contract.CustosialVaultMetaData.ABI)))
		if err != nil {
			return nil, fmt.Errorf("abi json error: %v", err)
		}

		var tokenCustodyEvent TokenCustodyEvent
		err = contractAbi.UnpackIntoInterface(&tokenCustodyEvent, tokenCustodyEventName, vLog.Data)
		if err != nil {
			return nil, err
		}
		tokenCustodyEvent.TokenId = big.NewInt(vLog.Topics[1].Big().Int64())

		return &tokenCustodyEvent, nil
	default:
		log.Warning("default case...")
		return nil, fmt.Errorf("unknown event signature hash: %v", vLog.Topics[0].Hex())
	}
}
