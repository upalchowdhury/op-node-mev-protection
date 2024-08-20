// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract MockEAS {
    // Struct to store attestations
    struct Attestation {
        string txHash;
        string encryptedData; // Encrypted private data
        address attester;     // Address of the entity who created the attestation
    }

    // Mapping from transaction hash to the attestation
    mapping(string => Attestation) public attestations;

    // Event emitted when an attestation is created
    event AttestationCreated(string txHash, address indexed attester);

    // Function to submit an attestation with encrypted private data
    function attestTransactionWithPrivateData(string memory txHash, string memory encryptedData) public {
        // Ensure the attestation does not already exist
        require(bytes(attestations[txHash].txHash).length == 0, "Attestation already exists");

        // Create the attestation
        attestations[txHash] = Attestation({
            txHash: txHash,
            encryptedData: encryptedData,
            attester: msg.sender
        });

        // Emit an event for the attestation creation
        emit AttestationCreated(txHash, msg.sender);
    }

    // Function to verify the attestation data
    function verifyAttestation(string memory txHash, string memory encryptedData) public view returns (bool) {
        // Ensure the attestation exists
        require(bytes(attestations[txHash].txHash).length != 0, "Attestation does not exist");

        // Verify that the encrypted data matches
        return keccak256(abi.encodePacked(attestations[txHash].encryptedData)) == keccak256(abi.encodePacked(encryptedData));
    }

    // Function to retrieve attestation data
    function getAttestation(string memory txHash) public view returns (string memory, string memory, address) {
        // Ensure the attestation exists
        require(bytes(attestations[txHash].txHash).length != 0, "Attestation does not exist");

        // Return the attestation data
        Attestation memory attestation = attestations[txHash];
        return (attestation.txHash, attestation.encryptedData, attestation.attester);
    }
}
