package gethrpcclient

import (
	"bytes"
	"encoding/json"
	"io"
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

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Id      int    `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Error   *Error
}

type StringResponse struct {
	*Response
	Result string
}

type DataResponse struct {
	*Response
	Result Data
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
	Result *Hash
}

type TransactionResponse struct {
	*Response
	Result *TransactionResult
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

	return bytes.NewReader(br)
}
