package containers

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum"
	"go.uber.org/dig"

	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/simple_storage"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/user_nft"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/voting"
)

type UserNFTAddress common.Address
type VotingContractAddress common.Address
type SimpleStorageAddress common.Address

func (c *Container) SolidityProvider() {
	// Provide *ethclient.Client directly
	if err := c.Container.Provide(func(cfg *config.AppConfig) (*ethclient.Client, error) {
		ethClient, err := ethereum.NewEthereumClient(cfg)
		if err != nil {
			log.Printf("Failed to initialize Ethereum client: %v", err)
			return nil, err
		}
		return ethClient.Eth, nil // Provide as *ethclient.Client
	}); err != nil {
		c.Error = err
	}

	// Provide *ethclient.Client as bind.ContractBackend
	if err := c.Container.Provide(func(ethClient *ethclient.Client) bind.ContractBackend {
		return ethClient // Provide as bind.ContractBackend
	}); err != nil {
		c.Error = err
	}

	// Provide UserNFTAddress
	if err := c.Container.Provide(func(cfg *config.AppConfig) UserNFTAddress {
		return UserNFTAddress(common.HexToAddress(cfg.UserNFTContractAddress))
	}); err != nil {
		c.Error = err
	}

	// Provide VotingContractAddress
	if err := c.Container.Provide(func(cfg *config.AppConfig) VotingContractAddress {
		return VotingContractAddress(common.HexToAddress(cfg.VotingContractAddress))
	}); err != nil {
		c.Error = err
	}

	// Provide SimpleStorageAddress if needed
	if err := c.Container.Provide(func(cfg *config.AppConfig) SimpleStorageAddress {
		return SimpleStorageAddress(common.HexToAddress(cfg.SimpleStorageContractAddress))
	}); err != nil {
		c.Error = err
	}

	// Provide *simple_storage.Simplestorage contract instance
	if err := c.Container.Provide(func(address SimpleStorageAddress, backend bind.ContractBackend) (*simple_storage.Simplestorage, error) {
		return simple_storage.NewSimplestorage(common.Address(address), backend)
	}); err != nil {
		c.Error = err
	}

	// Provide *user_nft.Solidity contract instance
	if err := c.Container.Provide(func(address UserNFTAddress, backend bind.ContractBackend) (*user_nft.Solidity, error) {
		return user_nft.NewSolidity(common.Address(address), backend)
	}); err != nil {
		c.Error = err
	}

	// Provide *voting.Voting contract instance
	if err := c.Container.Provide(func(address VotingContractAddress, backend bind.ContractBackend) (*voting.Voting, error) {
		return voting.NewVoting(common.Address(address), backend)
	}); err != nil {
		c.Error = err
	}

	// Provide transactor for Voting contract
	if err := c.Container.Provide(func(cfg *config.AppConfig) (*bind.TransactOpts, error) {
		auth, err := ethereum.NewTransactorWithPrivateKey(cfg.Ethereum.VotingPrivateKey, cfg.Ethereum.ChainID)
		if err != nil {
			log.Printf("Failed to create voting transactor: %v", err)
			return nil, err
		}
		return auth, nil
	}, dig.Name("VotingTransactor")); err != nil {
		c.Error = err
	}

	// Provide transactor for UserNFT contract
	if err := c.Container.Provide(func(cfg *config.AppConfig) (*bind.TransactOpts, error) {
		auth, err := ethereum.NewTransactorWithPrivateKey(cfg.Ethereum.UserNFTPrivateKey, cfg.Ethereum.ChainID)
		if err != nil {
			log.Printf("Failed to create user NFT transactor: %v", err)
			return nil, err
		}
		return auth, nil
	}, dig.Name("UserNFTTransactor")); err != nil {
		c.Error = err
	}
}
