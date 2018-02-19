package gethrpcclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

const ContentType = "application/json"

const (
	LatestBlock   = "latest"
	EarliestBlock = "earliest"
	PendingBlock  = "pending"
)

type Client struct {
	// Url for connection to geth
	Url string

	// Request id
	requestIterator RequestIterator
}

type requestIterator struct {
	current int
}

func (ri *requestIterator) Next() int {
	ri.current += 1

	return ri.current
}

func NewClient(url string) (c *Client) {
	c = &Client{
		Url:             url,
		requestIterator: &requestIterator{0},
	}

	return c
}

func (c *Client) newRequest(method string) (r *Request) {
	id := c.requestIterator.Next()
	fmt.Println(id, method)
	return &Request{
		Id:      id,
		JsonRPC: "2.0",
		Method:  method,
	}
}

// send rpc message
func (c *Client) send(request *Request, response RpcResponse) error {
	r, err := http.Post(c.Url, ContentType, request.jsonReader())
	if err != nil {
		panic(err)
	}

	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	return json.Unmarshal(buffer, response)
}
