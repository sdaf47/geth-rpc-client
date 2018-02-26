pragma solidity ^0.4.19;

contract Example {

    mapping (address => uint) balances;

    mapping (address => uint) debt;

    event Transfer(address, address, uint256);

    function balanceOf(address _holder) returns (uint256) {
        return balances[_holder];
    }

    function transfer(address _from, address _to, uint256 _amount) returns (bool) {
        balances[_to] += _amount;
        debt[_from] += _amount;
        Transfer(_from, _to, _amount);
        return true;
    }

    function debugTransfer(address _from, address _to, uint256 _amount) returns (bool) {
        balances[_to] += _amount;
        debt[_from] += _amount;
        Transfer(_from, _to, _amount);
        return true;
    }
}
