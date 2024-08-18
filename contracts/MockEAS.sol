// contracts/MockEAS.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MockEAS {
    mapping(string => bool) public attestations;

    function attestTransaction(string memory txHash) public {
        attestations[txHash] = true;
    }

    function verifyAttestation(string memory txHash) public view returns (bool) {
        return attestations[txHash];
    }
}
