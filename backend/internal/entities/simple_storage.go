package entities

import "math/big"

// SetNumberRequest represents the request structure for setting a number in SimpleStorage.
type SetNumberRequest struct {
	Value string `json:"value" validate:"required,numeric"`
}

// SetNumberResponse represents the response structure for the SetNumber endpoint.
type SetNumberResponse struct {
	TransactionHash string `json:"transactionHash"`
}

// GetNumberResponse represents the response structure for the GetNumber endpoint.
type GetNumberResponse struct {
	StoredValue *big.Int `json:"storedValue"`
}
