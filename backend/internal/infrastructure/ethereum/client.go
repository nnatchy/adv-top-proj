package ethereum

import (
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nnatchy/adv-top-proj/internal/config"
)

type Client struct {
	Eth *ethclient.Client
}

// NewEthereumClient initializes and returns a new Ethereum client.
func NewEthereumClient(cfg *config.AppConfig) (*Client, error) {
	rpcURL := cfg.Ethereum.RPCURL
	if rpcURL == "" {
		return nil, errors.New("ETHEREUM_RPC_URL environment variable is not set")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
		return nil, err
	}

	log.Println("Ethereum client initialized successfully")
	return &Client{Eth: client}, nil
}

// Backend returns the underlying ethclient.Client, which implements bind.ContractBackend.
func (c *Client) Backend() bind.ContractBackend {
	return c.Eth
}

// NewTransactorWithPrivateKey creates a new TransactOpts object for signing transactions with a given private key.
func NewTransactorWithPrivateKey(privateKey *ecdsa.PrivateKey, chainID int64) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return nil, err
	}
	return auth, nil
}
