package core_test

import (
	"context"
	"errors"
	"github.com/AlexRipoll/go-bridge/blockchain/core"
	"github.com/AlexRipoll/go-bridge/blockchain/core/evm"
	"github.com/golang/mock/gomock"
	"math/big"
	"testing"
)

func TestNewBridge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blockchainMock := evm.NewMockBlockchain(ctrl)
	t.Run("initialize new bridge instance", func(t *testing.T) {
		_, err := core.NewBridge(blockchainMock)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
}

func TestBridgeRetainNFT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	blockchainMock := evm.NewMockBlockchain(ctrl)

	b, err := core.NewBridge(blockchainMock)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	t.Run("successfully retain NFT and mint", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "ethereum"
		destination := "polygon"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm.Tx{
			Hash:      "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:      0,
			ChainID:   big.NewInt(1),
			Data:      []byte("data"),
			Gas:       300000,
			Nonce:     1,
			To:        wallet,
		}
		blockchainMock.EXPECT().RetainNFT(ctx, tokenId).Return(&evmTx, nil)
		blockchainMock.EXPECT().Mint(ctx, destination, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("successfully burn copy and release NFT", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "polygon"
		destination := "ethereum"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm.Tx{
			Hash:      "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:      0,
			ChainID:   big.NewInt(1),
			Data:      []byte("data"),
			Gas:       300000,
			Nonce:     1,
			To:        wallet,
		}
		blockchainMock.EXPECT().Burn(ctx, origin, tokenId).Return(&evmTx, nil)
		blockchainMock.EXPECT().ReleaseNFT(ctx, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("successfully burn copy and mint", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "polygon"
		destination := "binance"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm.Tx{
			Hash:      "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:      0,
			ChainID:   big.NewInt(1),
			Data:      []byte("data"),
			Gas:       300000,
			Nonce:     1,
			To:        wallet,
		}
		blockchainMock.EXPECT().Burn(ctx, origin, tokenId).Return(&evmTx, nil).Times(1)
		blockchainMock.EXPECT().Mint(ctx, destination, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("fail to transfer NFT", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "ethereum"
		destination := "polygon"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		blockchainMock.EXPECT().RetainNFT(ctx, tokenId).Return(nil, errors.New("tx reverted"))

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err == nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
}
