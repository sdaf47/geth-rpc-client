package gethrpcclient

const (
	NetVersion   = "net_version"
	NetPeerCount = "net_peerCount"
	NetListening = "net_listening"
)

// Returns the current network id.
func (c *Client) Version() (*NumberResponse, error) {
	request := c.newRequest(NetVersion)

	response := &NumberResponse{}

	return response, c.send(request, response)
}

// Returns number of peers currently connected to the client.
func (c *Client) PeerCount() (*QuantityResponse, error) {
	request := c.newRequest(NetPeerCount)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns true if client is actively listening for network connections.
func (c *Client) Listening() (*BooleanResponse, error) {
	request := c.newRequest(NetListening)

	response := &BooleanResponse{}

	return response, c.send(request, response)
}
