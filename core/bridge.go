package core
//
//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/AlexRipoll/go-bridge/core/event"
//	"github.com/AlexRipoll/go-bridge/core/evm"
//	"github.com/AlexRipoll/go-bridge/sys/storage"
//	"log"
//	"math/big"
//	"unsafe"
//)
//
//type Bridge struct {
//	mainNetwork string
//	custodian   evm.Custodian
//	bridgers    map[string]evm.Erc721Token
//	storage     storage.Storage
//}
//
//func NewBridge(mainNetwork string, vault evm.Custodian, bridger map[string]evm.Erc721Token) (*Bridge, error) {
//	return &Bridge{
//		mainNetwork: mainNetwork,
//		custodian:   vault,
//		bridgers:    bridger,
//	}, nil
//}
//
//type Tx struct {
//	Hash      string
//	Size      float64
//	ChainID   *big.Int
//	Gas       uint64
//	GasPrice  *big.Int
//	GasTipCap *big.Int
//	GasFeeCap *big.Int
//	Value     *big.Int
//	Nonce     uint64
//}
//
//type TxData struct {
//	Wallet      string   `json:"wallet"`
//	TokenId     *big.Int `json:"token_id"`
//	Origin      string   `json:"origin"`
//	Destination string   `json:"destination"`
//}
//
//// TransferNFT initiates the transfer of a token from one blockchain to another.
//// If from is the native blockchain, then the token is retained in the custody vault. Once the Tx event is received by
//// the Listener, it will be digested and the CompleteTransfer will be notified once the action needed can be
//// executed.
//// In case the token is transferred from a non-native blockchain to the native blockchain, the token copy will be burnt
//// and the token in the custody wallet will be released.
//// In case the token is transferred between non-native blockchains, the token copy from the origin will be burnt and
//// a copy of the token will be minted in the destination blockchain.
//func (b Bridge) TransferNFT(ctx context.Context, destination, origin, walletAddress string, tokenId *big.Int) error {
//	if origin == destination {
//		return errors.New("destination blockchain must be different than origin")
//	}
//
//	buf, err := json.Marshal(TxData{
//		Wallet:      walletAddress,
//		TokenId:     tokenId,
//		Origin:      origin,
//		Destination: destination,
//	})
//	if err != nil {
//		return err
//	}
//
//	if origin == b.mainNetwork {
//		tx, err := b.retainNFT(ctx, tokenId)
//		if err != nil {
//			return err
//		}
//
//		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
//			return err
//		}
//		log.Printf("retain nft: %v", tx)
//	} else if destination == b.mainNetwork {
//		tx, err := b.burnNFT(ctx, origin, tokenId)
//		if err != nil {
//			return err
//		}
//		log.Printf("burn nft: %#v", tx)
//
//		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
//			return err
//		}
//		log.Printf("retain nft: %v", tx)
//	} else {
//		tx, err := b.burnNFT(ctx, origin, tokenId)
//		if err != nil {
//			return err
//		}
//		log.Printf("burn nft: %#v", tx)
//
//		if err := b.storage.Put([]byte(tx.Hash), buf); err != nil {
//			return err
//		}
//		log.Printf("retain nft: %v", tx)
//	}
//
//	return nil
//}
//
//func (b Bridge) WalletTokens(address string, blockchain string) ([]*big.Int, error) {
//	if _, exists := b.bridgers[blockchain]; !exists {
//		return nil, fmt.Errorf("unknown client network %s", blockchain)
//	}
//	return b.bridgers[blockchain].TokensOf(address)
//}
//
//
//func (b Bridge) releaseNFT(ctx context.Context, walletAddress string, tokenId *big.Int) (*Tx, error) {
//	tx, err := b.custodian.ReleaseToken(ctx, walletAddress, tokenId)
//	if err != nil {
//		return nil, err
//	}
//
//	return toTx(tx), nil
//}
//
//func (b Bridge) mintNFT(ctx context.Context, destination, walletAddress string, tokenId *big.Int) (*Tx, error) {
//	bridger, ok := b.bridgers[destination]
//	if !ok {
//		return nil, errors.New("unknown destination blockchain")
//	}
//	tx, err := bridger.Mint(ctx, walletAddress, tokenId)
//	if err != nil {
//		return nil, err
//	}
//
//	return toTx(tx), nil
//}
//
//func (b Bridge) burnNFT(ctx context.Context, origin string, tokenId *big.Int) (*Tx, error) {
//	bridger, ok := b.bridgers[origin]
//	if !ok {
//		return nil, errors.New("unknown origin blockchain")
//	}
//	tx, err := bridger.Burn(ctx, tokenId)
//	if err != nil {
//		return nil, err
//	}
//
//	return toTx(tx), nil
//}
//
//func toTx(tx *evm.Tx) *Tx {
//	return (*Tx)(unsafe.Pointer(tx))
//}
