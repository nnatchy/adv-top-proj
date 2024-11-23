package repositories

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/simple_storage"
)

type ISimpleStorageRepository interface {
	SetNumber(ctx context.Context, value *big.Int, auth *bind.TransactOpts) (string, error)
	GetNumber(ctx context.Context) (*big.Int, error)
}

type SimpleStorageRepository struct {
	contract *simple_storage.Simplestorage // Generated contract binding
}

func NewSimpleStorageRepository(contract *simple_storage.Simplestorage) ISimpleStorageRepository {
	return &SimpleStorageRepository{
		contract: contract,
	}
}

// SetNumber sets a number in the SimpleStorage contract
func (r *SimpleStorageRepository) SetNumber(ctx context.Context, value *big.Int, auth *bind.TransactOpts) (string, error) {
	tx, err := r.contract.SetNumber(auth, value) // Call the SetNumber method on the contract
	if err != nil {
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// GetNumber retrieves the stored number from the SimpleStorage contract
func (r *SimpleStorageRepository) GetNumber(ctx context.Context) (*big.Int, error) {
	return r.contract.GetNumber(&bind.CallOpts{Context: ctx}) // Call the GetNumber method on the contract
}
