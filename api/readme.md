**Blockchain Project API**

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
