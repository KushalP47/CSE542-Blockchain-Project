# Blockchain Project API

This repository contains the source code for the API of a Blockchain project, which implements a basic blockchain application with peer-to-peer (P2P) networking capabilities using the libp2p library. It enables decentralized network interaction, transaction handling, and blockchain state management.

## Overview

The API consists of various handlers designed to manage accounts, transactions, blocks, and other blockchain-related operations. The project also includes essential components for P2P networking, block management, and a database for persistent data storage.

## Features

### **P2P Networking**

- **Root and Peer Nodes:** The application supports both root and peer nodes. Root nodes act as the initial point of contact in the network, while peer nodes connect to existing nodes using known addresses.
- **Custom Protocols:** The network uses custom protocols for node communication, including `PING`/`PONG` messages and transaction data sharing.
- **Peer Management:** Periodic updates ensure that each node maintains an active list of known peers and connections.

### **Transaction Handling**

- **Signed Transactions:** Nodes can send and receive signed transactions, which are validated and then added to the blockchain.
- **Transaction Management:** Includes handlers like `AddTxnHandler` to add transactions and `ReadTxnHandler` to retrieve transaction details.

### **Account Management**

- **Create and Manage Accounts:** The API provides functionalities to create (`AddAccountHandler`), retrieve (`GetAccountHandler`), and manage accounts on the blockchain.
- **Account Details:** Retrieves and displays account details, including the address, nonce, and balance.

### **Block Management**

- **Genesis Block Creation:** A dedicated handler, `Creategenesis.go`, is available to create the genesis block.
- **Block Retrieval:** Handlers like `Getblocknumberhandler.go`, `Getblockwithhandler.go`, and `Blockwithnumberhandler.go` allow for retrieving blocks by number or hash.

### **Database Integration**

The project uses the BadgerDB key-value store for managing blockchain data, including accounts, blocks, and transactions.

- **Account Data Management:** Functions in `account.go` read and write account data to the database, and calculate the state root hash.
- **Block Data Management:** `block.go` handles block data storage and retrieval, including fetching the last block from the database.
- **Transaction Data Management:** `transaction.go` manages transaction data, including reading, writing, and counting stored transactions.

## Components

### **API Handlers**

1. **AddAccountHandler** (`/add-account`, `POST`)

   - Adds a new account with a specified address and balance.
   - **Request Body:**
     ```json
     {
     	"address": "0x...",
     	"balance": 100
     }
     ```
   - **Response:**
     ```json
     {
     	"address": "0x...",
     	"nonce": 0,
     	"balance": 100
     }
     ```

2. **AddTxnHandler** (`/add-txn`, `POST`)

   - Adds a signed transaction to the blockchain.
   - **Request Body:**
     ```json
     {
     	"signed": "0x..."
     }
     ```

3. **GetAccountHandler** (`/get-account`, `POST`)

   - Retrieves account details.
   - **Request Body:**
     ```json
     {
     	"address": "0x..."
     }
     ```
   - **Response:**
     ```json
     {
     	"address": "0x...",
     	"nonce": 0,
     	"balance": 100
     }
     ```

4. **ReadTxnHandler** (`/read-txn`, `POST`)

   - Retrieves details of a transaction using its hash.
   - **Request Body:**
     ```json
     {
     	"hash": "0x..."
     }
     ```

5. **Creategenesis.go** (`/create-genesis-block`)

   - Creates a genesis block for the blockchain.

6. **Balances.go** (`/get-balance`, `POST`)

   - Retrieves the balance of an Ethereum address.
   - **Request Body:**
     ```json
     {
     	"address": "0x..."
     }
     ```
   - **Response:**
     ```json
     {
     	"balance": 100
     }
     ```

7. **Getblocknumberhandler.go** (`/get-block-number`)

   - Retrieves the latest block number and its hash.

8. **Getblockwithhandler.go** (`/get-block-with-hash`)

   - Retrieves a block using a specified hash.

9. **Blockwithnumberhandler.go** (`/get-block-with-number`)

   - Retrieves a block using a specified block number.

10. **Noncehandler.go** (`/get-nonce`)
    - Retrieves the nonce of an Ethereum address.

### **Core Files**

- **`main.go`**

  - The entry point of the application. It starts the HTTP server for API requests and initializes the P2P network. Depending on the arguments, it can start as either a root node or a peer node.

- **`rootNode.go`**

  - Manages the behavior of a root node, including handling PING messages and adding transactions to the blockchain.

- **`peerNode.go`**

  - Defines the behavior of peer nodes, which connect to the network and participate in P2P communication.

- **`protocol.go`**

  - Implements custom protocols for node-to-node communication, including message structures for sending and receiving data.

- **`message.go`**
  - Contains functions to send specific messages like `PING`, `PONG`, and signed transactions.

### **Database Files**

- **`db.go`**

  - Initializes and manages the database connection, providing methods for reading and writing data.

- **`account.go`**

  - Manages user accounts in the database, handling data retrieval and state root hash calculation.

- **`block.go`**

  - Handles block-related data storage and retrieval from the database.

- **`transaction.go`**
  - Manages transactions within the database, including reading, writing, and counting transactions.

## Getting Started

### Prerequisites

- Go 1.18 or higher
- libp2p library

### Running the Application

1. Clone the repository:
   ```bash
   git clone https://github.com/KushalP47/CSE542-Blockchain-Project.git
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build and run the API:
   ```bash
   go run main.go
   ```

## Dependencies

- [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum): Go client for Ethereum.
- [dgraph-io/badger](https://github.com/dgraph-io/badger): Key-value store database.
- [golang.org/x/crypto/sha3](https://pkg.go.dev/golang.org/x/crypto/sha3): SHA-3 cryptographic hash functions.
