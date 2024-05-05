**\# Blockchain Project API**

This repository contains the source code for the API of the Blockchain
project.

**Overview**

The API consists of several handlers for managing accounts and
transactions within the blockchain network. Each handler serves a
specific purpose and interacts with the underlying blockchain
implementation.

**Handlers**

1\. AddAccountHandler

\- Description: Adds a new account to the blockchain.

\- Endpoint: \`/add-account\`

\- Method: \`POST\`

\- Request Body:

{

"address": "0x...",

"balance": 100

}

\- \`address’: The address of the account.

\- \`balance\`: The initial balance of the account.

**Response**:

{

"address": "0x...",

"nonce": 0,

"balance": 100

}

\- \`address\`: The address of the created account.

\- \`nonce\`: The nonce of the account.

\- \`balance\`: The balance of the account.

**2. AddTxnHandler**

\- Description: Handles transaction requests, adds transactions to the
blockchain.

\- Endpoint: ‘/add-txn\`

\- Method: \`POST\`

\- Request Body:

{

"signed": "0x..."

}

\- \`signed\`: The signed transaction.

\-

**Response:**

Status code indicating success or failure.

**3. GetAccountHandler**

\- Description: Retrieves account details from the blockchain.

\- Endpoint: \`/get-account\`

\- Method: \`POST\`

\- Request Body:

{

"address": "0x..."

}

\- \`address\`: The address of the account to retrieve.

**Response:**

{

"address": "0x...",

"nonce": 0,

"balance": 100

}

\- \`address\`: The address of the account.

\- \`nonce\`: The nonce of the account.

\- \`balance\`: The balance of the account.

**4. ReadTxnHandler**

\- Description: Retrieves transaction details from the blockchain.

\- Endpoint: \`/read-txn\`

\- Method: \`POST\`

\- Request Body:

{

"hash": "0x..."

}

\- \`hash\`: The hash of the transaction to retrieve.

**Response:**

Status code indicating success or failure.

**5. Creategenesis.go:**

\- Endpoint: \`/create-genesis-block\`

\- Description: Creates a genesis block for the blockchain.

**6. Balances.go:**

\- Endpoint: \`/get-balance\`

\- Method: \`POST\`

\- Description: Retrieves the balance of a given Ethereum address.

\- Request Body:

{

"address": "0x..."

}

**- Response:**

{

"balance": 100

}

7\. **Getblocknumberhandler.go:**

\- Endpoint: \`/get-block-number\`

\- Description: Retrieves the latest block number and its hash.

8\. **Getblockwithhandler.go:**

\- Endpoint: \`/get-block-with-hash\`

\- Description: Retrieves a block with a specified hash.

9\. **Blockwithnumberhandler.go:**

\- Endpoint: \`/get-block-with-number\`

\- Description: Retrieves a block with a specified block number.

10\. **Noncehandler.go:**

\- Endpoint: \`/get-nonce\`

\- Description: Retrieves the nonce of a given Ethereum address.

11\. **Getaccounthandler.go:**

\- Endpoint: \`/get-account\`

\- Method: \`POST\`

\- Description: Retrieves account details from the blockchain.

\- Request Body:

{

"address": "0x..."

}

**- Response:**

{

"address": "0x...",

"nonce": 0,

"balance": 100

}

**Usage**

1\. Clone the repository: \`git clone
https://github.com/yourusername/blockchain-api.git\`

2\. Install dependencies: \`go mod tidy\`

3\. Build and run the API: \`go run main.go\`

**Dependencies**

\- \[ethereum/go-ethereum\](https://github.com/ethereum/go-ethereum): Go
client for Ethereum.

**\# Blockchain Project Blockchain**

This package holds the tools needed to create and manage a blockchain.

**Overview**

It handles blocks, transactions, and accounts within the blockchain.

**Components**

Block Management

\- Block: Represents a block in the chain, with key details.

\- Header: Stores metadata about a block.

**Transaction Management**

\- Txn: Holds transaction details.

\- SignedTx: A signed transaction.

\- VerifyTxn: Checks if a sender can make a transaction.

\- AddTxn: Adds a transaction to the blockchain.

**Account Management**

\- Account: Stores account info.

\- CreateAccount: Make a new account.

\- GetAccount: Retrieves an account.

\- UpdateAccount: Changes an account's balance.

**Usage**

1\. Import the \`blockchain\` package.

2\. Use its functions to interact with the blockchain.

3\. Ensure required dependencies like Ethereum libraries are set up.

**Dependencies**

\- ethereum/go-ethereum: Go Ethereum client.

\- dgraph-io/badger: Key-value store database.

\- golang.org/x/crypto/sha3: SHA-3 cryptographic hash functions.

**\# Blockchain Project Database**

**Database Package**

The \`database\` package contains functionalities related to managing
data in the blockchain database. It provides methods to interact with
the underlying database to read and write various data types such as
accounts, blocks, and transactions.

**Contents**

1\. \[db.go\]

2\. \[account.go\]

3\. \[block.go\]

4\. \[transaction.go\]

**db.go**

This file contains the \`Database\` struct and its associated methods.
It initializes the database, reads data, writes data, and closes the
database connection.

**account.go**

The \`account.go\` file contains functions related to managing user
accounts in the database. It provides methods to read and write account
data, as well as to calculate the state root hash.

**block.go**

The \`block.go\` file contains functions related to managing blocks in
the database. It provides methods to read and write block data, as well
as to retrieve the last block from the database.

**transaction.go**

The \`transaction.go\` file contains functions related to managing
transactions in the database. It provides methods to read and write
transaction data, as well as to retrieve transactions and count the
number of transactions stored in the database.
