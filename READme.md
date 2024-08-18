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
Transaction Simulation:

The node continuously simulates new transactions and applies a time-lock. These transactions are added to the sequencer's transaction pool with a specified unlock time.
EAS Attestation:

Before a transaction can be included in a block, it must be attested to by the EAS contract. The EASManager handles this attestation process.
Block Creation:

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
