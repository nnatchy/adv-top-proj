package handlers

import (
	"go.uber.org/dig"
)

// Handler ...
type Handler struct {
	Gateway Gateway
}

// Gateway ...
type Gateway struct {
	dig.In
	UserHandler          *UserHandler
	VotingHandler        *VotingHandler
	CandidateHandler     *CandidateHandler
}

// NewHandler ...
func NewHandler(g Gateway) *Handler {
	return &Handler{
		Gateway: g,
	}
}
