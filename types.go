package gethrpcclient

import (
	"math/big"
	"strconv"
)

const Decimal = "1000000000000000000"

// Some data as hex
type Data string

// 20 bytes
type Address Data

// 32 bytes
type Hash Data

// 8 bytes
type Nonce Data

// 256 bytes
type LongData Data

type ByteBoolean string

type Quantity string

type Number string

// compilers (solidity, serpent etc)
type Compiler string

type Filter struct {
	FromBlock Quantity // todo may be tag
	ToBlock   Quantity // todo may be tag
	Address   Address
	Topics    []Hash
}

type Param struct {
	Name string
	Type string
}

type AbiDefinition struct {
	constant bool
	Inputs   []Param
	Outputs  []Param
	Name     string
	Type     string
}

type ContractInfo struct {
	// contract source code
	Source string

	Language string

	LanguageVersion string

	CompilerVersion string

	AbiDefinition AbiDefinition

	UserDoc      interface{}
	DeveloperDoc interface{}
}

type Contract struct {
	Code string
	Info ContractInfo
}

// For filters created with c.NewFilter (eth_newFilter) logs are object
// todo https://github.com/ethereum/wiki/wiki/JSON-RPC#returns-42
type Log struct {
	Removed string // todo check

	// Integer of the log index position in the block. null when its pending log.
	LogIndex *Quantity

	// Integer of the transactions index position log was created from. null when its pending log.
	TransactionIndex Quantity

	// Hash of the transactions this log was created from. null when its pending log.
	TransactionHash Hash

	// Hash of the block where this log was in. null when its pending. null when its pending log.
	BlockHash *Hash

	// The block number where this log was in. null when its pending. null when its pending log.
	BlockNumber *Quantity

	// Address from which this log originated.
	Address Address

	// Contains one or more 32 Bytes non-indexed arguments of the log.
	Data Data

	// Array of 0 to 4 32 Bytes DATA of indexed log arguments.
	// (In solidity: The first topic is the hash of the signature of the event
	// (e.g. Deposit(address,bytes32,uint256)),
	// except you declared the event with the anonymous specifier.)
	Topics []Hash
}

type Block struct {
	// the block number. null when its pending block.
	Number *Quantity

	// hash of the block. null when its pending block.
	Hash *Hash

	// hash of the parent block.
	ParentHash Hash

	// hash of the generated proof-of-work. null when its pending block.
	Nonce *Nonce

	// SHA3 of the uncles data in the block.
	Sha3Uncles Hash

	// the bloom filter for the logs of the block. null when its pending block.
	LogsBloom *LongData

	// the root of the transaction trie of the block.
	TransactionsRoot Hash

	// the root of the final state trie of the block.
	StateRoot Hash

	//  the root of the receipts trie of the block.
	ReceiptsRoot Hash

	// the address of the beneficiary to whom the mining rewards were given.
	Miner Address

	// integer of the difficulty for this block.
	Difficulty Quantity

	// integer of the total difficulty of the chain until this block.
	TotalDifficulty Quantity

	// the "extra data" field of this block.
	ExtraData Data

	// integer the size of this block in bytes.
	Size Quantity

	// the maximum gas allowed in this block.
	GasLimit Quantity

	// the total used gas by all transactions in this block.
	GasUsed Quantity

	// the unix timestamp for when the block was collated.
	Timestamp Quantity

	// Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter.
	Transactions []Hash

	// Array of uncle hashes.
	Uncles []Hash
}

type TransactionReceipt struct {
	// hash of the transaction.
	TransactionHash Hash

	// integer of the transactions index position in the block.
	TransactionIndex Quantity

	// hash of the block where this transaction was in.
	BlockHash Hash

	// block number where this transaction was in.
	BlockNumber Quantity

	// The total amount of gas used when this transaction was executed in the block.
	CumulativeGasUsed Quantity

	// The amount of gas used by this specific transaction alone.
	GasUsed Quantity

	// The contract address created, if the transaction was a contract creation, otherwise null.
	ContractAddress *Address

	// Array of log objects, which this transaction generated.
	Logs []Log

	// Bloom filter for light clients to quickly retrieve related logs.
	LogsBloom LongData
}

func (n Number) Int() (int, error) {
	return strconv.Atoi(string(n))
}

// Convert hex to bigint
func (b Quantity) BigInt() *big.Int {
	balance := new(big.Int)
	balance.SetString(string(b)[2:], 16)

	return balance
}

// Convert hex to int
func (b Quantity) Int() int {
	// todo
	return 0
}

// Convert hex to bool
func (b ByteBoolean) Bool() bool {
	if b[len(b)-1] == '1' {
		return true
	}
	return false
}
