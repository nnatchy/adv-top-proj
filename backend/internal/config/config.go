package config

import (
	"crypto/ecdsa"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort                      string         `mapstructure:"APP_PORT"`
	AppEnv                       string         `mapstructure:"APP_ENV"`
	AllowOrigin                  AllowOrigin    `mapstructure:",squash"` // Include AllowOrigin struct
	Ethereum                     EthereumConfig `mapstructure:",squash"`
	Jwt                          JwtConfig      `mapstructure:",squash"`
	Database                     PostgresConfig `mapstructure:",squash"`
	UserNFTContractAddress       string         `mapstructure:"USER_NFT_CONTRACT_ADDRESS"`
	VotingContractAddress        string         `mapstructure:"VOTING_CONTRACT_ADDRESS"`
	SimpleStorageContractAddress string         `mapstructure:"SIMPLE_STORAGE_CONTRACT_ADDRESS"`
}

type AllowOrigin struct {
	FrontendURLs []string `mapstructure:"FRONTEND_URLS"` // Frontend URLs for CORS
}

type PostgresConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DbName   string `mapstructure:"POSTGRES_DB"`
}

type EthereumConfig struct {
	RPCURL            string            `mapstructure:"ETHEREUM_RPC_URL"`
	ApiKey            string            `mapstructure:"ETHEREUM_API_KEY"`
	ChainID           int64             `mapstructure:"ETHEREUM_CHAIN_ID"`
	VotingPrivateKey  *ecdsa.PrivateKey // For Voting contract interactions
	UserNFTPrivateKey *ecdsa.PrivateKey // For UserNFT contract interactions
}

type JwtConfig struct {
	JWT_SECRET_KEY string `mapstructure:"JWT_SECRET_KEY"`
}

// LoadConfig reads configuration from the environment and converts the private keys
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// Split FrontendURLs from a comma-separated string into a slice
	frontendURLs := viper.GetString("FRONTEND_URLS")
	config.AllowOrigin.FrontendURLs = strings.Split(frontendURLs, ",")

	// Load the private key for Voting contract
	votingPrivateKeyHex := viper.GetString("ETHEREUM_PRIVATE_KEY_VOTING")
	votingPrivateKey, err := crypto.HexToECDSA(votingPrivateKeyHex)
	if err != nil {
		log.Printf("Error converting voting private key, %s", err)
		return nil, err
	}
	config.Ethereum.VotingPrivateKey = votingPrivateKey

	// Load the private key for UserNFT contract
	userNFTPrivateKeyHex := viper.GetString("ETHEREUM_PRIVATE_KEY_USER_NFT")
	userNFTPrivateKey, err := crypto.HexToECDSA(userNFTPrivateKeyHex)
	if err != nil {
		log.Printf("Error converting user NFT private key, %s", err)
		return nil, err
	}
	config.Ethereum.UserNFTPrivateKey = userNFTPrivateKey

	return &config, nil
}
