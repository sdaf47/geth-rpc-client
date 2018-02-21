package gethrpcclient

import (
	"math/big"
	"strconv"
	"fmt"
	"encoding/json"
	"github.com/pkg/errors"
)

const Decimal = "1000000000000000000"

// Some data as hex
type Data string

func NewData(data []byte) *Data {
	d := Data(data)

	return &d
}

// 20 bytes
type Address Data

func NewAddress(addr Data) *Address {
	a := Address(addr)

	return &a
}

// 32 bytes
type Hash Data

// 8 bytes
type Nonce Data

// 256 bytes
type LongData Data

type ByteBoolean string

type Quantity struct {
	*big.Int
}

func NewQuantity(big *big.Int) *Quantity {
	return &Quantity{big}
}

type Number string

// compilers (solidity, serpent etc)
type Compiler string

type Filter struct {
	FromBlock Quantity `json:"fromBlock"`
	ToBlock   Quantity `json:"toBlock"`
	Address   Address  `json:"address"`
	Topics    []Hash   `json:"topics"`
}

type Param struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AbiDefinition struct {
	Constant bool    `json:"constant"`
	Inputs   []Param `json:"inputs"`
	Outputs  []Param `json:"outputs"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
}

type ContractInfo struct {
	// contract source code
	Source string `json:"source"`

	Language string `json:"language"`

	LanguageVersion string `json:"languageVersion"`

	CompilerVersion string `json:"compilerVersion"`

	AbiDefinition AbiDefinition `json:"abiDefinition"`

	UserDoc      interface{} `json:"userDoc"`
	DeveloperDoc interface{} `json:"developerDoc"`
}

type Contract struct {
	Code string       `json:"code"`
	Info ContractInfo `json:"info"`
}

// For filters created with c.NewFilter (eth_newFilter) logs are object
// todo https://github.com/ethereum/wiki/wiki/JSON-RPC#returns-42
type Log struct {
	Removed string `json:"removed"` // todo check

	// Integer of the log index position in the block. null when its pending log.
	LogIndex *Quantity `json:"logIndex"`

	// Integer of the transactions index position log was created from. null when its pending log.
	TransactionIndex Quantity `json:"transactionIndex"`

	// Hash of the transactions this log was created from. null when its pending log.
	TransactionHash Hash `json:"transactionHash"`

	// Hash of the block where this log was in. null when its pending. null when its pending log.
	BlockHash *Hash `json:"blockHash"`

	// The block number where this log was in. null when its pending. null when its pending log.
	BlockNumber *Quantity `json:"blockNumber"`

	// Address from which this log originated.
	Address Address `json:"address"`

	// Contains one or more 32 Bytes non-indexed arguments of the log.
	Data Data `json:"data"`

	// Array of 0 to 4 32 Bytes DATA of indexed log arguments.
	// (In solidity: The first topic is the hash of the signature of the event
	// (e.g. Deposit(address,bytes32,uint256)),
	// except you declared the event with the anonymous specifier.)
	Topics []Hash `json:"topics"`
}

type Block struct {
	// the block number. null when its pending block.
	Number *Quantity `json:"number"`

	// hash of the block. null when its pending block.
	Hash *Hash `json:"hash"`

	// hash of the parent block.
	ParentHash Hash `json:"parentHash"`

	// hash of the generated proof-of-work. null when its pending block.
	Nonce *Nonce `json:"nonce"`

	// SHA3 of the uncles data in the block.
	Sha3Uncles Hash `json:"sha3Uncles"`

	// the bloom filter for the logs of the block. null when its pending block.
	LogsBloom *LongData `json:"logsBloom"`

	// the root of the transaction trie of the block.
	TransactionsRoot Hash `json:"transactionsRoot"`

	// the root of the final state trie of the block.
	StateRoot Hash `json:"stateRoot"`

	//  the root of the receipts trie of the block.
	ReceiptsRoot Hash `json:"receiptsRoot"`

	// the address of the beneficiary to whom the mining rewards were given.
	Miner Address `json:"miner"`

	// integer of the difficulty for this block.
	Difficulty Quantity `json:"difficulty"`

	// integer of the total difficulty of the chain until this block.
	TotalDifficulty Quantity `json:"totalDifficulty"`

	// the "extra data" field of this block.
	ExtraData Data `json:"extraData"`

	// integer the size of this block in bytes.
	Size Quantity `json:"size"`

	// the maximum gas allowed in this block.
	GasLimit Quantity `json:"gasLimit"`

	// the total used gas by all transactions in this block.
	GasUsed Quantity `json:"gasUsed"`

	// the unix timestamp for when the block was collated.
	Timestamp Quantity `json:"timestamp"`

	// Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter.
	Transactions []Hash `json:"transactions"`

	// Array of uncle hashes.
	Uncles []Hash `json:"uncles"`
}

type TransactionReceipt struct {
	// hash of the transaction.
	TransactionHash Hash `json:"transactionHash"`

	// integer of the transactions index position in the block.
	TransactionIndex Quantity `json:"transactionIndex"`

	// hash of the block where this transaction was in.
	BlockHash Hash `json:"blockHash"`

	// block number where this transaction was in.
	BlockNumber Quantity `json:"blockNumber"`

	// The total amount of gas used when this transaction was executed in the block.
	CumulativeGasUsed Quantity `json:"cumulativeGasUsed"`

	// The amount of gas used by this specific transaction alone.
	GasUsed Quantity `json:"gasUsed"`

	// The contract address created, if the transaction was a contract creation, otherwise null.
	ContractAddress *Address `json:"contractAddress"`

	// Array of log objects, which this transaction generated.
	Logs []Log `json:"logs"`

	// Bloom filter for light clients to quickly retrieve related logs.
	LogsBloom LongData `json:"logsBloom"`
}

func (n Number) Int() (int, error) {
	return strconv.Atoi(string(n))
}

func (q *Quantity) MarshalJSON() (d []byte, err error) {
	if q.Int == nil {
		return json.Marshal(nil)
	}

	q16 := fmt.Sprintf("%x", q)

	if string(q16[0]) == "0" {
		q16 = q16[1:]
	}
	q16 = "0x" + q16

	return json.Marshal(q16)
}

func (q *Quantity) UnmarshalJSON(d []byte) (err error) {
	var data string

	if err = json.Unmarshal(d, &data); err != nil {
		return err
	}

	quantity := new(big.Int)
	quantity, ok := quantity.SetString(data[2:], 16)
	if !ok {
		return errors.New("Can`t set string: " + data[2:])
	}

	q.Int = quantity

	return nil
}

// Convert hex to bigint
//func (b Quantity) BigInt() *big.Int {
//	balance := new(big.Int)
//	balance.SetString(string(b)[2:], 16)
//
//	return balance
//}

// Convert hex to bool
func (b ByteBoolean) Bool() bool {
	if b[len(b)-1] == '1' {
		return true
	}
	return false
}

// The ethereum transaction call object
type Transaction struct {
	// 20 Bytes - (optional) The address the transaction is sent from.
	From *Address `json:"from"`

	// 20 Bytes - The address the transaction is directed to.
	To *Address `json:"to"`

	// (optional) Integer of the gas provided for the transaction execution.
	// eth_call consumes zero gas, but this parameter may be needed
	// by some executions.
	Gas *Quantity `json:"gas"`

	// (optional) Integer of the gasPrice used for each paid gas
	GasPrice *Quantity `json:"gas_price"`

	// (optional) Integer of the value send with this transaction
	Value *Quantity `json:"value"`

	// (optional) Hash of the method signature and encoded parameters.
	// For details see Ethereum Contract ABI
	Data Data `json:"data"`

	// (optional) Integer of a nonce. This allows to overwrite your own pending
	// transactions that use the same nonce.
	Nonce *Quantity `json:"nonce"`
}

type TransactionResult struct {
	*Transaction

	// hash of the transaction.
	Hash *Hash `json:"hash"`

	// hash of the block where this transaction was in. null when its pending.
	BlockHash *Hash `json:"blockHash"`

	// block number where this transaction was in. null when its pending.
	BlockNumber *Quantity `json:"blockNumber"`

	// integer of the transactions index position in the block. null when its pending.
	TransactionIndex *Quantity `json:"transactionIndex"`

	//  the data send along with the transaction
	Input *Data `json:"input"`
}
