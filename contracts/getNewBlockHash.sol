// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

contract CheckBlockHash {

    uint256 private lastBlockTimestamp;
    uint256 private secondLastBlockTimestamp;

    function storeCurrentTimestamp() public {
        secondLastBlockTimestamp = lastBlockTimestamp;
        lastBlockTimestamp = block.timestamp;
    }

    function getTimestampDifference() public view returns (uint256) {
        return lastBlockTimestamp - secondLastBlockTimestamp;
    }

    // Return the block hash of a given block number
    function currentBlockHash() public view returns (bytes32) {
        return blockhash(block.number);
    }

    // Return the block hash of the previous block
    function getPreviousBlockHash() public view returns (bytes32) {
        return blockhash(block.number - 1);
    }
}