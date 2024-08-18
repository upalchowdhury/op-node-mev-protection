# Optimistic Rollup Node with EAS Time-Locking and MEV Protection


This project implements a simplified optimistic rollup node in Go, featuring Ethereum Attestation Service (EAS) integration for transaction attestations and time-locking to mitigate MEV (Miner Extractable Value) attacks. The node interacts with a local Ethereum blockchain (Anvil via Foundry), simulates transactions with time-locks, and verifies them through the EAS before adding them to blocks.



### Architecture Overview
The architecture revolves around three main components:

Rollup Node: The core logic responsible for handling transactions, creating blocks, and interacting with L1 (Anvil) and the EAS contract.
Sequencer: Manages the transaction pool and ensures that only time-unlocked and validly attested transactions are included in the blocks.
EAS Manager: Interacts with a deployed Ethereum Attestation Service contract to attest and verify transactions before inclusion in a block.
### Key Components
Anvil: A local Ethereum testnet that simulates a blockchain environment.
Mock EAS Contract: A simple contract deployed on Anvil to simulate transaction attestations.
Go Rollup Node: Implements the rollup node behavior, including time-locking transactions and verifying them with EAS.

### Process Flow
- Transaction Simulation:

The node continuously simulates new transactions and applies a time-lock. These transactions are added to the sequencer's transaction pool with a specified unlock time.
- EAS Attestation:

Before a transaction can be included in a block, it must be attested to by the EAS contract. The EASManager handles this attestation process.
- Block Creation:

The sequencer creates blocks from transactions that have passed their unlock time and have been successfully attested by the EAS. These blocks are then validated and added to the chain.
Monitoring:

The node continuously monitors the creation and validation of blocks and logs important events, such as transactions being unlocked and blocks being created.



### Features
- Time-Locked Transactions: Transactions are locked until a specified unlock time, mitigating MEV attacks.
- EAS Integration: Transactions are attested and verified using the Ethereum Attestation Service before being included in a block.
- Block Creation: Transactions that pass their time-lock and attestation are bundled into blocks that are validated and added to the chain.



### Steps to run
- Install foundry
- Run local node using `anvil`
- Update `run_test.sh` script's `ANVIL_ACCOUNT` and `ANVIL_PRIVATE_KEY` with `account` and `private key` from anvil output
- Run `run_test.sh`


### Future Improvements
- MEV Protection Enhancements: Implement additional MEV protection strategies, such as encrypted transaction pools.

- Optimistic Rollup Features: Expand the functionality to include more features from full optimistic rollup implementations, such as fraud proofs.

- Extended EAS Functionality: Integrate more advanced attestation features from the Ethereum Attestation Service.

- Execution Verification by Sequencer
The rollup node allows users to query whether a transaction has been executed by the sequencer. This can be verified using a simple query to the node’s execution log.

- Private Data in Attestations
Attestations can include private data that is encrypted before being attested on-chain. This data can only be accessed by authorized recipients with the correct decryption keys.

- Off-Chain Attestations
Off-chain attestations differ from simple signatures in that they involve a structured proof process that can be stored off-chain (e.g., IPFS). These attestations are signed with private keys to provide authenticity and integrity, similar to on-chain attestations.

- Merkle Root and Private Data
When transactions are sequentially processed, the Merkle root is computed at the end of the batch. To maintain the relevance of private data, its hash is committed as part of the transaction data, ensuring that it contributes to the overall batch integrity.

- User Notifications for Attestations
The node includes an event listener that triggers notifications to users when an attestation is created with them as a recipient. Notifications can be delivered via webhooks, email, or other messaging systems.

