package gethrpcclient

const (
	EthAccounts                            = "eth_accounts"
	EthGetBalance                          = "eth_getBalance"
	EthGetCode                             = "eth_getCode"
	EthCall                                = "eth_call"
	EthTransaction                         = "eth_sendTransaction"
	EthProtocolVersion                     = "eth_protocolVersion"
	EthSyncing                             = "eth_syncing"
	EthCoinbase                            = "eth_coinbase"
	EthMining                              = "eth_mining"
	EthHashrate                            = "eth_hashrate"
	EthGasPrice                            = "eth_gasPrice"
	EthBlockNumber                         = "eth_blockNumber"
	EthGetStorageAt                        = "eth_getStorageAt"
	EthGetTransactionCount                 = "eth_getTransactionCount"
	EthGetBlockTransactionCountByHash      = "eth_getBlockTransactionCountByHash"
	EthGetBlockTransactionCountByNumber    = "eth_getBlockTransactionCountByNumber"
	EthGetUncleCountByBlockHash            = "eth_getUncleCountByBlockHash"
	EthGetUncleCountByBlockNumber          = "eth_getUncleCountByBlockNumber"
	EthSign                                = "eth_sign"
	EthSendRawTransaction                  = "eth_sendRawTransaction"
	EthEstimateGas                         = "eth_estimateGas"
	EthGetBlockByHash                      = "eth_getBlockByHash"
	EthGetBlockByNumber                    = "eth_getBlockByNumber"
	EthGetTransactionByHash                = "eth_getTransactionByHash"
	EthGetTransactionByBlockHashAndIndex   = "eth_getTransactionByBlockHashAndIndex"
	EthGetTransactionByBlockNumberAndIndex = "eth_getTransactionByBlockNumberAndIndex"
	EthGetTransactionReceipt               = "eth_getTransactionReceipt"
	EthGetUncleByBlockHashAndIndex         = "eth_getUncleByBlockHashAndIndex"
	EthGetUncleByBlockNumberAndIndex       = "eth_getUncleByBlockNumberAndIndex"
	EthNewFilter                           = "eth_newFilter"
	EthNewBlockFilter                      = "eth_newBlockFilter"
	EthNewPendingTransactionFilter         = "eth_newPendingTransactionFilter"
	EthUninstallFilter                     = "eth_uninstallFilter"
	EthGetFilterChanges                    = "eth_getFilterChanges"
	EthGetFilterLogs                       = "eth_getFilterLogs"
	EthGetLogs                             = "eth_getLogs"
	EthGetWork                             = "eth_getWork"
	EthSubmitWork                          = "eth_submitWork"
	EthSubmitHashrate                      = "eth_submitHashrate"
)

// Returns accounts with keys
func (c *Client) Accounts() (*AddressesResponse, error) {
	request := c.newRequest(EthAccounts)

	response := &AddressesResponse{}

	return response, c.send(request, response)
}

// Returns the balance of the account of given address
func (c *Client) GetBalance(addr Address, block string) (*QuantityResponse, error) {
	request := c.newRequest(EthGetBalance)

	request.Params = []string{
		string(addr),
		block,
	}

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns code at a given address
func (c *Client) GetCode(addr Address, block string) (*StringResponse, error) {
	request := c.newRequest(EthGetCode)

	request.Params = []string{
		string(addr),
		block,
	}

	response := &StringResponse{}

	return response, c.send(request, response)
}

// New message call immediately without creating a transaction on the block chain
func (c *Client) Call(tx *Transaction, block string, response RpcResponse) error {
	request := c.newRequest(EthCall)

	request.Params = []interface{}{
		tx,
		block,
	}

	return c.send(request, response)
}

// New message call with creating a transaction on the block chain
func (c *Client) SendTransaction(tx *Transaction) (*HashResponse, error) {
	request := c.newRequest(EthTransaction)

	request.Params = []interface{}{
		tx,
	}

	response := &HashResponse{}

	return response, c.send(request, response)
}

// Creates new message call transaction or a contract creation for signed transactions.
func (c *Client) SendRowTransaction(tx Data) (*HashResponse, error) {
	request := c.newRequest(EthSendRawTransaction)

	request.Params = []interface{}{
		tx,
	}

	response := &HashResponse{}

	return response, c.send(request, response)
}

// Returns the current ethereum protocol version.
func (c *Client) ProtocolVersion() (*NumberResponse, error) {
	request := c.newRequest(EthProtocolVersion)

	response := &NumberResponse{}

	return response, c.send(request, response)
}

// Returns an object with data about the sync status or false.
func (c *Client) Syncing() (*SyncingResponse, error) {
	request := c.newRequest(EthSyncing)

	response := &SyncingResponse{}

	return response, c.send(request, response)
}

// Returns the client coinbase address.
func (c *Client) Coinbase() (*AddressResponse, error) {
	request := c.newRequest(EthCoinbase)

	response := &AddressResponse{}

	return response, c.send(request, response)
}

// Returns the client coinbase address.
func (c *Client) Mining() (*BooleanResponse, error) {
	request := c.newRequest(EthMining)

	response := &BooleanResponse{}

	return response, c.send(request, response)
}

// Returns the number of hashes per second that the node is mining with.
func (c *Client) HashRate() (*QuantityResponse, error) {
	request := c.newRequest(EthHashrate)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the current price per gas in wei.
func (c *Client) GasPrice() (*QuantityResponse, error) {
	request := c.newRequest(EthGasPrice)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of most recent block.
func (c *Client) BlockNumber() (*QuantityResponse, error) {
	request := c.newRequest(EthBlockNumber)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the value from a storage position at a given address.
func (c *Client) GetStorageAt(addr Address, index Quantity, block string) (*DataResponse, error) {
	request := c.newRequest(EthGetStorageAt)

	request.Params = []interface{}{
		addr,
		index,
		block,
	}
	response := &DataResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions sent from an address.
func (c *Client) GetTransactionCount(addr Address, block string) (*QuantityResponse, error) {
	request := c.newRequest(EthGetTransactionCount)

	request.Params = []string{
		string(addr),
		block,
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (c *Client) GetBlockTransactionCountByHash(addr Hash) (*QuantityResponse, error) {
	request := c.newRequest(EthGetBlockTransactionCountByHash)

	request.Params = []string{
		string(addr),
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (c *Client) GetBlockTransactionCountByNumber(block string) (*QuantityResponse, error) {
	request := c.newRequest(EthGetBlockTransactionCountByNumber)

	request.Params = []string{
		string(block),
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (c *Client) GetUncleCountByBlockHash(block string) (*QuantityResponse, error) {
	request := c.newRequest(EthGetUncleCountByBlockHash)

	request.Params = []string{
		string(block),
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (c *Client) GetUncleCountByBlockNumber(block string) (*QuantityResponse, error) {
	request := c.newRequest(EthGetUncleCountByBlockNumber)

	request.Params = []string{
		string(block),
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// he sign method calculates an Ethereum specific signature with:
// sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))).
// The address to sign with must be unlocked.
func (c *Client) Sign(addr Address, message Data) (*DataResponse, error) {
	request := c.newRequest(EthSign)

	request.Params = []string{
		string(addr),
		string(message),
	}
	response := &DataResponse{}

	return response, c.send(request, response)
}

// Generates and returns an estimate of how much gas is necessary
// to allow the transaction to complete.
// The transaction will not be added to the blockchain.
// Note that the estimate may be significantly more than the amount of
// gas actually used by the transaction, for a variety of reasons
// including EVM mechanics and node performance.
// todo optional parameters
func (c *Client) EstimateGas(tx *Transaction, block string) (*QuantityResponse, error) {
	request := c.newRequest(EthEstimateGas)

	request.Params = []interface{}{
		tx,
		block,
	}

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (c *Client) GetBlockByHash(block string, full bool) (*BlockResponse, error) {
	request := c.newRequest(EthGetBlockByHash)

	request.Params = []interface{}{
		string(block),
		full,
	}
	response := &BlockResponse{}

	return response, c.send(request, response)
}

// Returns information about a block by block number.
func (c *Client) GetBlockByNumber(block string, full bool) (*BlockResponse, error) {
	request := c.newRequest(EthGetBlockByNumber)

	request.Params = []interface{}{
		string(block),
		full,
	}
	response := &BlockResponse{}

	return response, c.send(request, response)
}

// Returns the information about a transaction requested by transaction hash.
func (c *Client) GetTransactionByHash(tx Hash) (*TransactionResponse, error) {
	request := c.newRequest(EthGetTransactionByHash)

	request.Params = []string{
		string(tx),
	}
	response := &TransactionResponse{}

	return response, c.send(request, response)
}

// Returns information about a transaction by block hash and transaction index position.
func (c *Client) GetTransactionByBlockHashAndIndex(block Hash, index Quantity) (*TransactionResponse, error) {
	request := c.newRequest(EthGetTransactionByBlockHashAndIndex)

	request.Params = []interface{}{
		block,
		index,
	}
	response := &TransactionResponse{}

	return response, c.send(request, response)
}

// Returns information about a transaction by block number and transaction index position.
func (c *Client) GetTransactionByBlockNumberAndIndex(block string, index Quantity) (*TransactionResponse, error) {
	request := c.newRequest(EthGetTransactionByBlockNumberAndIndex)

	request.Params = []interface{}{
		block,
		index,
	}
	response := &TransactionResponse{}

	return response, c.send(request, response)
}

// Returns the receipt of a transaction by transaction hash.
func (c *Client) GetTransactionReceipt(tx Hash) (*TransactionReceiptResponse, error) {
	request := c.newRequest(EthGetTransactionReceipt)

	request.Params = []string{
		string(tx),
	}
	response := &TransactionReceiptResponse{}

	return response, c.send(request, response)
}

// Returns the receipt of a transaction by transaction hash.
func (c *Client) GetUncleByBlockHashAndIndex(block Hash, index Quantity) (*BlockResponse, error) {
	request := c.newRequest(EthGetUncleByBlockHashAndIndex)

	request.Params = []interface{}{
		block,
		index,
	}
	response := &BlockResponse{}

	return response, c.send(request, response)
}

// Returns information about a uncle of a block by number and uncle index position.
func (c *Client) GetUncleByBlockNumberAndIndex(block string, index Quantity) (*BlockResponse, error) {
	request := c.newRequest(EthGetUncleByBlockNumberAndIndex)

	request.Params = []interface{}{
		block,
		index,
	}
	response := &BlockResponse{}

	return response, c.send(request, response)
}

// Creates a filter object, based on filter options,
// to notify when the state changes (logs).
// To check if the state has changed, call GetFilterChanges (eth_getFilterChanges).
func (c *Client) NewFilter(filter Filter) (*QuantityResponse, error) {
	request := c.newRequest(EthNewFilter)

	request.Params = []interface{}{
		filter,
	}
	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Creates a filter in the node,
// to notify when a new block arrives.
// To check if the state has changed, call GetFilterChanges (eth_getFilterChanges).
func (c *Client) NewBlockFilter() (*QuantityResponse, error) {
	request := c.newRequest(EthNewBlockFilter)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Creates a filter in the node,
// to notify when new pending transactions arrive.
// To check if the state has changed, call GetFilterChanges (eth_getFilterChanges).
func (c *Client) NewPendingTransactionFilter() (*QuantityResponse, error) {
	request := c.newRequest(EthNewPendingTransactionFilter)

	response := &QuantityResponse{}

	return response, c.send(request, response)
}

// Uninstalls a filter with given id.
// Should always be called when watch is no longer needed.
// Additonally Filters timeout when they aren't requested with eth_getFilterChanges for a period of time.
func (c *Client) UninstallFilter(index Quantity) (*BooleanResponse, error) {
	request := c.newRequest(EthUninstallFilter)

	request.Params = []interface{}{
		index,
	}
	response := &BooleanResponse{}

	return response, c.send(request, response)
}

// Polling method for a filter, which returns an array of logs which occurred since last poll.
// todo may return 3 types of data
func (c *Client) GetFilterChanges(index Quantity) (*LogsResponse, error) {
	request := c.newRequest(EthGetFilterChanges)

	request.Params = []interface{}{
		index,
	}
	response := &LogsResponse{}

	return response, c.send(request, response)
}

// Returns an array of all logs matching filter with given id.
func (c *Client) GetFilterLogs(index Quantity) (*LogsResponse, error) {
	request := c.newRequest(EthGetFilterLogs)

	request.Params = []interface{}{
		index,
	}
	response := &LogsResponse{}

	return response, c.send(request, response)
}

// Returns an array of all logs matching filter with given id.
func (c *Client) GetLogs(filter *Filter) (*LogsResponse, error) {
	request := c.newRequest(EthGetLogs)

	request.Params = []interface{}{
		filter,
	}
	response := &LogsResponse{}

	return response, c.send(request, response)
}

// Returns the hash of the current block, the seedHash, and the boundary condition to be met ("target").
// Array - Array with the following properties:
//  - DATA, 32 Bytes - current block header pow-hash
//  - DATA, 32 Bytes - the seed hash used for the DAG.
//  - DATA, 32 Bytes - the boundary condition ("target"), 2^256 / difficulty.
func (c *Client) GetWork() (*HashesResponse, error) {
	request := c.newRequest(EthGetWork)

	response := &HashesResponse{}

	return response, c.send(request, response)
}

// Used for submitting a proof-of-work solution.
//  - DATA, 8 Bytes - The nonce found (64 bits)
//  - DATA, 32 Bytes - The header's pow-hash (256 bits)
//  - DATA, 32 Bytes - The mix digest (256 bits)
func (c *Client) SubmitWork(nonce Nonce, headers Hash, digest Hash) (*BooleanResponse, error) {
	request := c.newRequest(EthSubmitWork)

	request.Params = []string{
		string(nonce),
		string(headers),
		string(digest),
	}
	response := &BooleanResponse{}

	return response, c.send(request, response)
}

// Used for submitting mining hashrate.
//  - Hashrate, a hexadecimal string representation (32 bytes) of the hash rate
//  - String - A random hexadecimal(32 bytes) ID identifying the client
func (c *Client) SubmitHashrate(hashrate Hash, id Hash) (*BooleanResponse, error) {
	request := c.newRequest(EthSubmitHashrate)

	request.Params = []string{
		string(hashrate),
		string(id),
	}
	response := &BooleanResponse{}

	return response, c.send(request, response)
}
