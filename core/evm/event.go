package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type TokenCustodyEvent struct {
	TokenId     *big.Int
	Holder      common.Address
	Destination *big.Int
}

var tokenCustodyEventName = "TokenCustody"
var tokenCustodySig = []byte(fmt.Sprintf("%s(uint256,address,uint256)", tokenCustodyEventName))
var tokenCustodySigHash = crypto.Keccak256Hash(tokenCustodySig)
