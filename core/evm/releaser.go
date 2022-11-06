package evm

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/AlexRipoll/go-bridge/sys/storage"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type Releaser struct {
	Clients map[uint64]Client
	db storage.Db
}

func (r Releaser) releaseToken(ctx context.Context, rx event.Rx) error {
	log.Infof("initializing token release process to chain id %v...", rx.Destination.Int64())

	if err := r.waitForFinality(ctx, rx); err != nil {
		log.Error(fmt.Sprintf("wait for finality error (ethereum): %v", err.Error()))
		return err
	}

	key := make([]byte, len(rx.Origin.Bytes())+len(rx.TxHash))
	key = append(rx.Origin.Bytes())
	key = append(key, rx.TxHash[:]...)
	if err := r.db.Delete(key); err != nil {
		return err
	}

	return nil
}

func (r Releaser) waitForFinality(ctx context.Context, rx event.Rx) error {
	originClient, exists := r.Clients[rx.Origin.Uint64()]
	if !exists {
		return fmt.Errorf("no client with chain id %v in releaser", rx.Origin.Uint64())
	}
	destinationClient, exists := r.Clients[rx.Destination.Uint64()]
	if !exists {
		return fmt.Errorf("no client with chain id %v in releaser", rx.Destination.Uint64())
	}

	exists, err := destinationClient.contracts.erc721Token.Exists(ctx, rx.TokenId)
	if err != nil {
		return err
	}
	log.Infof("block amount for reaching finality: %v", destinationClient.finality)

	for {
		log.Info("waiting for reaching finality...")
		currentBlock, err := originClient.Transactor.CurrentBlock(ctx)
		if err != nil {
			return err
		}
		log.Infof("current block number: %v", currentBlock)
		log.Infof("TX block number: %v", rx.TxBlock)
		log.Infof("block amount left for reaching finality: %v", currentBlock-rx.TxBlock)
		log.Info("----------------------------------------------")
		log.Info(currentBlock - rx.TxBlock)
		log.Info((currentBlock - rx.TxBlock) >= destinationClient.Finality())
		log.Info("----------------------------------------------")

		if (currentBlock - rx.TxBlock) >= destinationClient.Finality() {
			log.Infof("finality reached")
			if !exists {
				if _, err := destinationClient.contracts.erc721Token.Mint(ctx, rx.Holder, rx.TokenId); err != nil {
					return err
				}
				log.Infof("minting token with Id %v and releasing it to address %v", rx.TokenId, rx.Holder)
				return nil
			}
			// check if owner is custodial vault
			if err := checkRetention(ctx, destinationClient, rx.TokenId); err != nil {
				return err
			}

			// updates the token owner in the holdCustody map if token has been transferred to another wallet
			if err := updateOwner(ctx, destinationClient, rx.TokenId, rx.Holder); err != nil {
				return err
			}

			time.Sleep(time.Second)
			_, err = destinationClient.contracts.custodianVault.ReleaseToken(ctx, rx.Holder, rx.TokenId)
			if err != nil {
				return err
			}
			log.Infof("token with Id %v released to address %v", rx.TokenId, rx.Holder)

			return nil
		}
		time.Sleep(time.Millisecond * 500)
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

func updateOwner(ctx context.Context, client Client, tokenId *big.Int, sender string) error {
	holder, err := client.contracts.custodianVault.TokenHolder(ctx, tokenId)
	if err != nil {
		log.Info("error fetching token %v holder: ", tokenId)
		return err
	}

	if holder == sender {
		return nil
	}

	tx, err := client.contracts.custodianVault.UpdateOwner(ctx, tokenId, sender)
	if err != nil {
		log.Info("error updating token %v ownership: ", tokenId)
		return err
	}
	log.Infof("token %v holder updated from %v to %v", tokenId, holder, sender)
	log.Infof("tx %#v", tx)

	return nil
}
