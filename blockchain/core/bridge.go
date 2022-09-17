package core

import (
	"context"
	"errors"
	"github.com/AlexRipoll/go-bridge/blockchain/core/evm"
	"log"
	"math/big"
	"unsafe"
)

type Bridge struct {
	custodian evm.Custodian
	bridgers map[string]evm.Bridger
}

func NewBridge(vault evm.Custodian, bridger map[string]evm.Bridger) (*Bridge, error) {
	return &Bridge{
		custodian: vault,
		bridgers: bridger,
	}, nil
}

type Tx struct {
	Hash string
	Size float64
	//From string
	ChainID   *big.Int
	Data      []byte
	Gas       uint64
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int
	Nonce     uint64
	To        string
}

type Rx struct {
	hash string
}

// TransferNFT transfers the specified token with tokenId from the origin blockchain to the destination blockchain.
// If from is the native blockchain, then the token is retained in the custody vault and a copy of it will be minted
// in the destination blockchain.
// In case the token is transferred from a non-native blockchain to the native blockchain, the token copy will be burnt
// and the token in the custody wallet will be released.
// In case the token is transferred between non-native blockchains, the token copy from the origin will be burnt and
// a copy of the token will be minted in the destination blockchain.
func (b Bridge) TransferNFT(ctx context.Context, destination, origin, walletAddress string, tokenId *big.Int) error {
	if origin == destination {
		return errors.New("destination blockchain must be different than origin")
	}
	if origin == "ethereum" {
		tx, err := b.retainNFT(ctx, tokenId)
		if err != nil {
			return err
		}
		log.Printf("retain nft: %#v", tx)
		// TODO add scanner to scan if the tx has been mined and then proceed with minting/burning/releasing
		tx, err = b.mintNFT(ctx, destination, walletAddress, tokenId)
		if err != nil {
			return err
		}
		log.Printf("mint nft: %#v", tx)
	} else if destination == "ethereum" {
		tx, err := b.burnNFT(ctx, origin, tokenId)
		if err != nil {
			return err
		}
		log.Printf("burn nft: %#v", tx)
		tx, err = b.releaseNFT(ctx, walletAddress, tokenId)
		if err != nil {
			return err
		}
		log.Printf("release nft: %#v", tx)
	} else {
		tx, err := b.burnNFT(ctx, origin, tokenId)
		if err != nil {
			return err
		}
		log.Printf("burn nft: %#v", tx)
		tx, err = b.mintNFT(ctx, destination, walletAddress, tokenId)
		if err != nil {
			return err
		}
		log.Printf("mint nft: %#v", tx)
	}

	return nil
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
	tx, err := bridger.Mint(ctx, destination, walletAddress, tokenId)
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
	tx, err := bridger.Burn(ctx, origin, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func toTx(tx *evm.Tx) *Tx {
	return (*Tx)(unsafe.Pointer(tx))
}
