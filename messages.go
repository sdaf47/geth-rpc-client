package gethrpcclient

import (
	"bytes"
	"encoding/json"
	"io"
	"fmt"
)

type RequestIterator interface {
	Next() int
}

type RpcResponse interface {
}

type Request struct {
	Id      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	JsonRPC string      `json:"jsonrpc"`
}

type Response struct {
	Id      int    `json:"id"`
	JsonRPC string `json:"jsonrpc"`
}

type StringResponse struct {
	*Response
	Result string
}

type DataResponse struct {
	*Response
	Result Data
}

type ContractResponse struct {
	*Response
	Result Contract
}

type BlockResponse struct {
	*Response
	Result *Block
}

type AddressesResponse struct {
	*Response
	Result []Address
}

type HashesResponse struct {
	*Response
	Result []Hash
}

type QuantityResponse struct {
	*Response
	Result Quantity
}

type NumberResponse struct {
	*Response
	Result Number
}

type SyncingResponse struct {
	*Response
	Result interface{} // todo
}

type AddressResponse struct {
	*Response
	Result Address
}

type BooleanResponse struct {
	*Response
	Result bool
}

type ByteBooleanResponse struct {
	*Response
	Result ByteBoolean
}

type HashResponse struct {
	*Response
	Result Hash
}

type TransactionResponse struct {
	*Response
	Result *Transaction
}

type TransactionReceiptResponse struct {
	*Response
	Result *TransactionReceipt
}

type CompilersResponse struct {
	*Response
	Result []Compiler
}

type LogsResponse struct {
	*Response
	Result []Log
}

func (r *Request) jsonReader() io.Reader {
	br, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(br))

	return bytes.NewReader(br)
}
