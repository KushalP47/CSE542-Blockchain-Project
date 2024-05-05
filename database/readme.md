**Blockchain Project Database**

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
