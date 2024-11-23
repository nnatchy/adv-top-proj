package repositories

import (
	"go.uber.org/dig"
)

// Repository ...
type Repository struct {
	RepositoryGateway RepositoryGateway
}

// RepositoryGateway ...
type RepositoryGateway struct {
	dig.In
	IUserRepository

	UserRepository          IUserRepository
	SimpleStorageRepository ISimpleStorageRepository
	VoteRepository          IVoteRepository
	CandidateRepository     ICandidateRepository
}

// NewRepository ...
func NewRepository(rg RepositoryGateway) *Repository {
	return &Repository{
		RepositoryGateway: rg,
	}
}
