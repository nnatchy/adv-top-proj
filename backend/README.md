# Backend: NFT Authentication and Ethereum Voting System

This is a **backend service** for handling NFT-based authentication and Ethereum-powered voting. It is built using **Golang** with the **Echo** framework and runs on port `8010`.

---

## Features

- **NFT Authentication**:
  - Users are authenticated using their NFTs.
  - Secure and transparent verification of ownership.

- **Ethereum Voting**:
  - Voting logic powered by Ethereum smart contracts.
  - Ensures immutability and trust in the voting process.

---

## Prerequisites

- **Golang** (version 1.18 or higher)
- **Ethereum Node/Provider** (e.g., Infura or Alchemy)
- Environment file (`.env`)

---

## Getting Started

Follow these steps to set up and run the project:

### 1. Clone the Repository
```bash
git clone <repository-url>
cd <repository-name>
```

### 2. Add the .env File
Place the required .env file in the root directory of the repository. The .env file should include the following variables:
```env
APP_PORT=8010
APP_ENV=dev
JWT_SECRET_KEY=
FRONTEND_URL=

# Ethereum Configuration
ETHEREUM_RPC_URL=
ETHEREUM_API_KEY=
ETHEREUM_CHAIN_ID=
ETHEREUM_PRIVATE_KEY_VOTING=
ETHEREUM_PRIVATE_KEY_USER_NFT=

# PostgreSQL Database Configuration
POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=

# Smart Contracts
USER_NFT_CONTRACT_ADDRESS=
VOTING_CONTRACT_ADDRESS=
```

### 3. Install Dependencies
Run the following command to install all required Go modules:
```bash
go mod tidy
```

### 4. Run the Server
```bash
go run cmd/server/main.go
```

The server will be available at:

```arduino
http://localhost:8010
```

## Run Using Docker
You can also use Docker to build and run the project.

### 1. Build the Docker Image
Run the following command to build the Docker image:
```bash
docker compose build
```

### 2. Run the Application
Start the application using Docker Compose:
```bash
docker compose up
```

### 3. Stop the Application
To stop the Docker container, run:
```bash
docker compose down
```

## Project Structure
```bash
├── cmd
│   └── server
│       └── main.go        # Main entry point for the server
├── internal
│   ├── config             # Configuration files and utilities
│   ├── constant           # Project-wide constant values
│   ├── containers         # Dependency injection and container management
│   ├── entities           # Data structures and models for business logic
│   ├── handlers           # API handlers (routes and HTTP logic)
│   ├── infrastructure     # Low-level implementations (e.g., database connections)
│   ├── middlewares        # Middleware for authentication, logging, etc.
│   ├── models             # Database models
│   ├── repositories       # Data access logic (interact with the database)
│   ├── services           # Core business logic (Ethereum and NFT operations)
│   └── utils              # Helper functions and utilities
├── go.mod                 # Go Module File
├── go.sum                 # Dependency Checksum
└── .env                   # Environment Variables (Excluded from Git)

```

## Technologies Used
  - **Golang Echo**: Web framework for building APIs.
  - **Ethereum**: Blockchain platform for handling voting logic.
  - **NFT**: Used for authentication based on token ownership.
  - **PostgreSQL**: Relational database for storing user and voting data.
