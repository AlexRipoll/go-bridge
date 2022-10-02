package core_test

import (
	"context"
	"errors"
	"github.com/AlexRipoll/go-bridge/core"
	evm2 "github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/golang/mock/gomock"
	"math/big"
	"testing"
)

func TestNewBridge(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	custodianMock := evm2.NewMockCustodian(ctrl)
	bridgers := make(map[string]evm2.Bridger)

	t.Run("initialize new bridge instance", func(t *testing.T) {
		_, err := core.NewBridge(custodianMock, bridgers)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
}

func TestBridgeRetainNFT(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	custodianMock := evm2.NewMockCustodian(ctrl)

	polygonBridgerMock := evm2.NewMockBridger(ctrl)
	hederaBridgerMock := evm2.NewMockBridger(ctrl)

	bridgers := make(map[string]evm2.Bridger)
	bridgers["polygon"] = polygonBridgerMock
	bridgers["hedera"] = hederaBridgerMock

	b, err := core.NewBridge(custodianMock, bridgers)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	t.Run("successfully retain Bridger and mint", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "ethereum"
		destination := "polygon"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm2.Tx{
			Hash:    "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:    0,
			ChainID: big.NewInt(1),
			Gas:     300000,
			Nonce:   1,
		}
		custodianMock.EXPECT().RetainNFT(ctx, tokenId).Return(&evmTx, nil)
		polygonBridgerMock.EXPECT().Mint(ctx, destination, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("successfully burn copy and release Bridger", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "polygon"
		destination := "ethereum"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm2.Tx{
			Hash:    "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:    0,
			ChainID: big.NewInt(1),
			Gas:     300000,
			Nonce:   1,
		}
		polygonBridgerMock.EXPECT().Burn(ctx, origin, tokenId).Return(&evmTx, nil)
		custodianMock.EXPECT().ReleaseNFT(ctx, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("successfully burn copy and mint", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "polygon"
		destination := "hedera"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm2.Tx{
			Hash:    "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:    0,
			ChainID: big.NewInt(1),
			Gas:     300000,
			Nonce:   1,
		}
		polygonBridgerMock.EXPECT().Burn(ctx, origin, tokenId).Return(&evmTx, nil)
		hederaBridgerMock.EXPECT().Mint(ctx, destination, wallet, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("fail to transfer Bridger", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "ethereum"
		destination := "polygon"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		custodianMock.EXPECT().RetainNFT(ctx, tokenId).Return(nil, errors.New("tx reverted"))

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err == nil {
			t.Fatalf("expected nil error, got %v", err)
		}
	})
	t.Run("unknown bridger blockchain", func(t *testing.T) {
		ctx := context.Background()
		tokenId := big.NewInt(0)
		origin := "ethereum"
		destination := "binance"
		wallet := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		evmTx := evm2.Tx{
			Hash:    "3a6eb0790f39ac87c94f3856b2dd2c5d110e6811602261a9a923d3bb23adc8b7",
			Size:    0,
			ChainID: big.NewInt(1),
			Gas:     300000,
			Nonce:   1,
		}

		custodianMock.EXPECT().RetainNFT(ctx, tokenId).Return(&evmTx, nil)

		if err := b.TransferNFT(ctx, destination, origin, wallet, tokenId); err == nil {
			t.Fatal("expected error, got nil")
		}
	})

}
