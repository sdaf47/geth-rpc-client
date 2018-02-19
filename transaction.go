package gethrpcclient

import (
	"math/big"
	"fmt"
)

// The ethereum transaction call object
type Transaction struct {
	// 20 Bytes - (optional) The address the transaction is sent from.
	From Address `json:"from"`

	// 20 Bytes - The address the transaction is directed to.
	To Address `json:"to"`

	// (optional) Integer of the gas provided for the transaction execution.
	// eth_call consumes zero gas, but this parameter may be needed
	// by some executions.
	Gas Quantity `json:"gas"`

	// (optional) Integer of the gasPrice used for each paid gas
	GasPrice Quantity `json:"gas_price"`

	// (optional) Integer of the value send with this transaction
	Value Quantity `json:"value"`

	// (optional) Hash of the method signature and encoded parameters.
	// For details see Ethereum Contract ABI
	Data *string `json:"data"`

	// (optional) Integer of a nonce. This allows to overwrite your own pending
	// transactions that use the same nonce.
	Nonce *Quantity `json:"nonce"`

	// the following properties will only be present when you receive transaction information
	// hash of the transaction.
	hash *Hash

	// hash of the block where this transaction was in. null when its pending.
	blockHash *Hash

	// block number where this transaction was in. null when its pending.
	blockNumber *Quantity

	// integer of the transactions index position in the block. null when its pending.
	transactionIndex *Quantity

	//  the data send along with the transaction
	input *Data
}

type TransactionBuilder interface {
	From(Address) TransactionBuilder
	To(Address) TransactionBuilder
	Value(*big.Int) TransactionBuilder
	Gas(*big.Int) TransactionBuilder
	GasPrice(*big.Int) TransactionBuilder
	Data(string) TransactionBuilder
	Nonce(*Quantity) TransactionBuilder
	Build() *Transaction
}

type transactionBuilder struct {
	t *Transaction
}

func NewTransactionBuilder() TransactionBuilder {
	tb := &transactionBuilder{}

	tb.t = &Transaction{}

	return tb
}

func (tb *transactionBuilder) From(addr Address) (TransactionBuilder) {
	tb.t.From = addr

	return tb
}

func (tb *transactionBuilder) To(to Address) (TransactionBuilder) {
	tb.t.To = to

	return tb
}

func (tb *transactionBuilder) Value(value *big.Int) (TransactionBuilder) {
	tb.t.Value = Quantity(fmt.Sprintf("%#x", value.Bytes()))

	return tb
}

func (tb *transactionBuilder) Gas(amount *big.Int) (TransactionBuilder) {
	tb.t.Gas = Quantity(fmt.Sprintf("%#x", amount.Bytes()))

	return tb
}

func (tb *transactionBuilder) GasPrice(price *big.Int) (TransactionBuilder) {
	tb.t.GasPrice = Quantity(fmt.Sprintf("%#x", price.Bytes()))

	return tb
}

func (tb *transactionBuilder) Data(data string) (TransactionBuilder) {
	tb.t.Data = &data

	return tb
}

func (tb *transactionBuilder) Nonce(id *Quantity) (TransactionBuilder) {
	tb.t.Nonce = id

	return tb
}

func (tb *transactionBuilder) Build() (*Transaction) {
	return tb.t
}
