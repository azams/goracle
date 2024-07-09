// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

contract Price {
    address payable public owner;

    constructor() {
        owner = payable(msg.sender);
    }

    modifier onlyOwner {
        require(owner == payable(msg.sender), "You are not an owner!");
        _;
    }

    function getContractAddress() public view returns (address) {
        return address(this);
    }

    struct Transaction {
        string coin;
        uint8 done;
    }

    struct Coin {
        string coin;
        string price;
    }

    Transaction[] private transactions;
    mapping(string=> string) coinPrice;
    Coin[] private coins;
    event NewTransaction(string coin);

    function requestUpdate(string memory _coin) public {
        Transaction memory newTransaction = Transaction(_coin, 0);
        transactions.push(newTransaction);
        // subimt event
        emit NewTransaction(_coin);
    }

    function getCoinPrice(string memory _coin) public view returns(string memory) {
        return coinPrice[_coin];
    }

    function updateCoinPrice(string memory _coin, string memory _price) public {
        require(owner == payable(msg.sender), "You are not the owner!");
        coinPrice[_coin] = _price;
    }

    function updateOwner(address newAddress) public onlyOwner {
        owner = payable(newAddress);
    }
}
