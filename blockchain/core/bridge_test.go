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
	t.Run("successfully retain NFT", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)

		evmTx := evm.Tx{
			Hash:      "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:      0,
			ChainID:   big.NewInt(1),
			Data:      []byte("data"),
			Gas:       300000,
			Nonce:     1,
			To:        "0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
		}
		blockchainMock.EXPECT().RetainNFT(ctx, tokenId).Return(&evmTx, nil)

		tx, err := b.RetainNFT(ctx, tokenId)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if tx == nil {
			t.Fatal("expected Tx, got nil")
		}
	})
	t.Run("successfully retain NFT", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)

		blockchainMock.EXPECT().RetainNFT(ctx, tokenId).Return(nil, errors.New("tx reverted"))

		tx, err := b.RetainNFT(ctx, tokenId)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if tx != nil {
			t.Fatalf("expected nil tx, got %v", tx)
		}
	})
}
