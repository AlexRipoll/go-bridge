package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type NftCustodyEvent struct {
	TokenId *big.Int
	Holder	common.Address
}

var nftCustodySig = []byte("NFTCustody(uint256,address)")
var nftCustodySigHash = crypto.Keccak256Hash(nftCustodySig)
