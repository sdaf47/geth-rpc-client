# geth-rpc-client
My own geth-client with blackjack and contracts

---

# Tests

For test use the ethereum private net:
```
$ geth --networkid 1234 --ipcpath ~/Library/Ethereum/geth.ipc --datadir testnet --rpc --rpcport 7545 --rpcapi "eth,net,web3,personal" console

```

You need two unlocked accounts with balance and mining.

## Example contract

```javascript
contract Example {

    mapping (address => uint) balances;
    mapping (address => uint) debt;

    event Transfer(address from, address to, uint amount);

    function transfer(address _from, address _to, uint _amount) {
        balances[_to] += _amount;
        debt[_from] += _amount;
        Transfer(_from, _to, _amount);
    }

    function debugTransfer(address _from, address _to, uint _amount) returns (bool _result) {
        balances[_to] += _amount;
        debt[_from] += _amount;
        Transfer(_from, _to, _amount);
        _result = true;
    }
}
```