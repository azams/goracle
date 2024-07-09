// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

contract RandomNumber {
    address payable public owner;

    constructor() {
        owner = payable(msg.sender);
    }

    modifier onlyOwner {
        require(owner == msg.sender, "You are not the owner!");
        _;
    }

    function getContractAddress() public view returns (address) {
        return address(this);
    }

    struct RandomRequest {
        uint256 id; // An ID for the request
        uint8 done; // Status: 0 for pending, 1 for completed
    }

    RandomRequest[] private requests;
    mapping(uint256 => uint256) public randomNumberMapping; // Map request IDs to random numbers
    event NewRandomRequest(uint256 requestId);

    function requestRandomNumber() public {
        uint256 newRequestId = requests.length;
        RandomRequest memory newRequest = RandomRequest(newRequestId, 0);
        requests.push(newRequest);
        emit NewRandomRequest(newRequestId);
    }

    function getRandomNumber() public view returns(uint256) {
        uint256 requestId = requests.length - 1;
        return randomNumberMapping[requestId];
    }

    function setRandomNumber(uint256 requestId, uint256 _randomNumber) public onlyOwner {
        require(requests[requestId].done == 0, "Request already completed");
        randomNumberMapping[requestId] = _randomNumber;
        requests[requestId].done = 1;
    }

    function updateOwner(address newAddress) public onlyOwner {
        owner = payable(newAddress);
    }
}
