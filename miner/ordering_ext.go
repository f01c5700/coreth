package miner

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/f01c5700/coreth/core/txpool"
	"github.com/f01c5700/coreth/core/types"
)

type TransactionsByPriceAndNonce = transactionsByPriceAndNonce

func NewTransactionsByPriceAndNonce(signer types.Signer, txs map[common.Address][]*txpool.LazyTransaction, baseFee *big.Int) *TransactionsByPriceAndNonce {
	return newTransactionsByPriceAndNonce(signer, txs, baseFee)
}
