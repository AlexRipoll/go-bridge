package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type NftCustodyEvent struct {
	TokenId *big.Int
	Holder  common.Address
}

var nftCustodyEventName = "NFTCustody"
var nftCustodySig = []byte(fmt.Sprintf("%s(uint256,address)", nftCustodyEventName))
var nftCustodySigHash = crypto.Keccak256Hash(nftCustodySig)
