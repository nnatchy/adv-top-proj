package services

import (
	"go.uber.org/dig"
)

// Service ...
type Service struct {
	ServiceGateway ServiceGateway
}

// ServiceGateway ...
type ServiceGateway struct {
	dig.In
	UserService          IUserService
	VoteService          IVoteService
	CandidateService     ICandidateService
}

// NewService ...
func NewService(sg ServiceGateway) *Service {
	return &Service{
		ServiceGateway: sg,
	}
}
