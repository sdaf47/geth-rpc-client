pragma solidity ^0.4.18;


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
