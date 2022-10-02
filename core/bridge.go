package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/AlexRipoll/go-bridge/sys/storage"
	"log"
	"math/big"
	"unsafe"
)

type Bridge struct {
	mainNetwork string
	custodian   evm.Custodian
	bridgers    map[string]evm.Bridger
	storage     storage.Storage
}

func NewBridge(mainNetwork string, vault evm.Custodian, bridger map[string]evm.Bridger) (*Bridge, error) {
	return &Bridge{
		mainNetwork: mainNetwork,
		custodian:   vault,
		bridgers:    bridger,
	}, nil
}

type Tx struct {
	Hash      string
	Size      float64
	ChainID   *big.Int
	Gas       uint64
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int
	Nonce     uint64
}

type TxData struct {
	Wallet      string   `json:"wallet"`
	TokenId     *big.Int `json:"token_id"`
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
}

// TransferNFT initiates the transfer of a token from one blockchain to another.
// If from is the native blockchain, then the token is retained in the custody vault. Once the Tx event is received by
// the Listener, it will be digested and the CompleteTransfer will be notified once the action needed can be
// executed.
// In case the token is transferred from a non-native blockchain to the native blockchain, the token copy will be burnt
// and the token in the custody wallet will be released.
// In case the token is transferred between non-native blockchains, the token copy from the origin will be burnt and
// a copy of the token will be minted in the destination blockchain.
func (b Bridge) TransferNFT(ctx context.Context, destination, origin, walletAddress string, tokenId *big.Int) error {
	if origin == destination {
		return errors.New("destination blockchain must be different than origin")
	}

	buf, err := json.Marshal(TxData{
		Wallet:      walletAddress,
		TokenId:     tokenId,
		Origin:      origin,
		Destination: destination,
	})
	if err != nil {
		return err
	}

	if origin == b.mainNetwork {
		tx, err := b.retainNFT(ctx, tokenId)
		if err != nil {
			return err
		}

		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
			return err
		}
		log.Printf("retain nft: %v", tx)
	} else if destination == b.mainNetwork {
		tx, err := b.burnNFT(ctx, origin, tokenId)
		if err != nil {
			return err
		}
		log.Printf("burn nft: %#v", tx)

		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
			return err
		}
		log.Printf("retain nft: %v", tx)
	} else {
		tx, err := b.burnNFT(ctx, origin, tokenId)
		if err != nil {
			return err
		}
		log.Printf("burn nft: %#v", tx)

		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
			return err
		}
		log.Printf("retain nft: %v", tx)
	}

	return nil
}

func (b Bridge) WalletTokens(address string, blockchain string) ([]*big.Int, error) {
	if _, exists := b.bridgers[blockchain]; !exists {
		return nil, fmt.Errorf("unknown client network %s", blockchain)
	}
	return b.bridgers[blockchain].TokensOf(address)
}

func (b Bridge) retainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	tx, err := b.custodian.RetainNFT(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func (b Bridge) releaseNFT(ctx context.Context, walletAddress string, tokenId *big.Int) (*Tx, error) {
	tx, err := b.custodian.ReleaseNFT(ctx, walletAddress, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func (b Bridge) mintNFT(ctx context.Context, destination, walletAddress string, tokenId *big.Int) (*Tx, error) {
	bridger, ok := b.bridgers[destination]
	if !ok {
		return nil, errors.New("unknown destination blockchain")
	}
	tx, err := bridger.Mint(ctx, walletAddress, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func (b Bridge) burnNFT(ctx context.Context, origin string, tokenId *big.Int) (*Tx, error) {
	bridger, ok := b.bridgers[origin]
	if !ok {
		return nil, errors.New("unknown origin blockchain")
	}
	tx, err := bridger.Burn(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func (b Bridge) Deploy(ctx context.Context) error {
	_, err := b.custodian.Deploy(ctx)
	if err != nil {
		return err
	}
	for _, bridger := range b.bridgers {
		_, err = bridger.Deploy(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b Bridge) CompleteTransfer(ctx context.Context, ch chan event.Rx) error {
	for {
		eventRx := <-ch

		buf, err := b.storage.Get([]byte(eventRx.TxHash.String()))
		if err != nil {
			return err
		}

		var txData TxData
		if err := json.Unmarshal(buf, &txData); err != nil {
			return err
		}

		switch eventRx.Action {
		case event.MintAction:
			tx, err := b.bridgers[txData.Destination].Mint(ctx, txData.Wallet, txData.TokenId)
			if err != nil {
				return err
			}
			log.Println("Mint Tx: ", tx)
		case event.ReleaseAction:
			tx, err := b.custodian.ReleaseNFT(ctx, txData.Wallet, txData.TokenId)
			if err != nil {
				return err
			}
			log.Println("Release Tx: ", tx)
		default:
			return fmt.Errorf("unknow action; %v", eventRx.Action)
		}
	}
}

func toTx(tx *evm.Tx) *Tx {
	return (*Tx)(unsafe.Pointer(tx))
}
