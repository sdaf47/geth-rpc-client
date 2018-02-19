package gethrpcclient

const (
	MethodPersonalNewAccount    = "personal_newAccount"
	MethodPersonalUnlockAccount = "personal_unlockAccount"
	MethodPersonalLockAccount   = "personal_lockAccount"
)

func (c *Client) NewAccount(pass string) (*AddressResponse, error) {
	request := c.newRequest(MethodPersonalNewAccount)

	request.Params = []string{pass}

	response := &AddressResponse{}

	return response, c.send(request, response)
}

func (c *Client) UnlockAccount(addr Address, pass string, duration int) (*BooleanResponse, error) {
	request := c.newRequest(MethodPersonalUnlockAccount)

	request.Params = []interface{}{
		addr,
		pass,
		duration,
	}

	response := &BooleanResponse{}

	return response, c.send(request, response)
}

func (c *Client) LockAccount(addr Address) (*BooleanResponse, error) {
	request := c.newRequest(MethodPersonalLockAccount)

	request.Params = []interface{}{
		addr,
	}

	response := &BooleanResponse{}

	return response, c.send(request, response)
}
