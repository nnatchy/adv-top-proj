package entities

import "github.com/golang-jwt/jwt"

type Claims struct {
	WalletAddress string `json:"walletAddress"`
	jwt.StandardClaims
}

type GetHealthUserRequest struct {
	Service string `json:"service"`
}

type GetHealthUserResponse struct {
	Status string `json:"status"`
}

type GetUserInfoByWalletResponse struct {
	ID             string `json:"id"`
	WalletAddress  string `json:"wallet_address"`
	Username       string `json:"username"`
	IsActive       bool   `json:"is_active"`
	StudentID      string `json:"student_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	HasVoteYear    bool   `json:"has_vote_year"`
	HasVoteRep     bool   `json:"has_vote_rep"`
}

type UpdateUserStatusRequest struct {
	HasVoteYear *bool `json:"has_vote_year,omitempty"`
	HasVoteRep  *bool `json:"has_vote_rep,omitempty"`
}