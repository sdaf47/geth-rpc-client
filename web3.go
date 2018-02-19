package gethrpcclient

import "fmt"

const (
	MethodWeb3Sha3          = "web3_sha3"
	MethodWeb3ClientVersion = "web3_clientVersion"
)

func (c *Client) Sha3(line string) (*DataResponse, error) {
	request := c.newRequest(MethodWeb3Sha3)

	param := fmt.Sprintf("%#x", []byte(line))
	request.Params = []string{
		param,
	}

	response := &DataResponse{}

	return response, c.send(request, response)
}

func (c *Client) ClientVersion() (*StringResponse, error) {
	request := c.newRequest(MethodWeb3ClientVersion)

	response := &StringResponse{}

	return response, c.send(request, response)
}
