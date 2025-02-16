# Mini Wallet API

This is a simple API backend service for managing a mini wallet, built with Go, Gin, GORM, and PostgreSQL. It implements endpoints for wallet initialization, enabling/disabling, depositing, withdrawing, and viewing transactions as specified in the Postman collection.

## Features

- **Initialize Wallet**: Create a wallet and obtain an API token.
- **Enable Wallet**: Activate the wallet for transactions.
- **View Wallet**: Check the wallet balance.
- **View Transactions**: Retrieve a list of all wallet transactions.
- **Deposit Money**: Add virtual money to the wallet.
- **Withdraw Money**: Withdraw funds from the wallet.
- **Disable Wallet**: Disable the wallet to stop transactions.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- Git (optional, for cloning the repository)

## Installation

### 1. Clone the Repository

Clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd <repository-directory>
```

### 2. Copy the Environment File

```bash
cp .env.example .env
```

### 3. Set Up the Database

```bash
psql -U postgres -c "CREATE DATABASE mini_wallet"
```

Update the database connection and credentials in the `.env` file.

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Application

```bash
go run main.go
```
