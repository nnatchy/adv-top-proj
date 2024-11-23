package services

import (
	"context"
	"errors"
	"time"

	"github.com/nnatchy/adv-top-proj/internal/config"
	"github.com/nnatchy/adv-top-proj/internal/entities"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/ethereum/solidity/user_nft"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"github.com/nnatchy/adv-top-proj/internal/repositories"
	"gorm.io/gorm"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.uber.org/dig"
)

type IUserService interface {
	PerformHealthCheck(context.Context) (*entities.GetHealthUserResponse, error)
	Mint(context.Context, string) error
	BalanceOf(context.Context, string) (int64, error)
	Burn(context.Context, string) error
	Login(ctx context.Context, walletAddress string) (*models.User, string, error)
	Logout(ctx context.Context) error
	Register(ctx context.Context, walletAddress, username, firstName, lastName, studentID string) error
	IsActivate(ctx context.Context, walletAddress string) (bool, error)
	GetUserInfoByWallet(ctx context.Context, walletAddress string) (entities.GetUserInfoByWalletResponse, error)
	UpdateUserStatus(ctx context.Context, walletAddress string, req *entities.UpdateUserStatusRequest) error
	AdminLogin(ctx context.Context, username, password string) (*models.Admin, error)
	CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error)
	GetCandidateByID(ctx context.Context, id string) (*models.Candidate, error)
	UpdateCandidate(ctx context.Context, id string, candidate *models.Candidate) (*models.Candidate, error)
	DeleteCandidate(ctx context.Context, id string) error
	GetCandidates(ctx context.Context, position string) ([]*models.Candidate, error)
}

type UserService struct {
	repo            repositories.RepositoryGateway
	log             *log.Logger
	userNFTContract *user_nft.Solidity
	transactor      *bind.TransactOpts
	config          *config.AppConfig
}

type UserServiceParams struct {
	dig.In
	Repo            repositories.RepositoryGateway
	Log             *log.Logger
	UserNFTContract *user_nft.Solidity
	Transactor      *bind.TransactOpts `name:"UserNFTTransactor"`
	Config          *config.AppConfig
}

func NewUserService(params UserServiceParams) IUserService {
	return &UserService{
		repo:            params.Repo,
		log:             params.Log,
		userNFTContract: params.UserNFTContract,
		transactor:      params.Transactor,
		config:          params.Config,
	}
}

func (s *UserService) PerformHealthCheck(ctx context.Context) (*entities.GetHealthUserResponse, error) {
	s.log.InfoWithID(ctx, "[Service: PerformHealthCheck] Performing health check on user-service")
	status, err := s.repo.CheckHealth(ctx)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: PerformHealthCheck] Failed to perform health check")
		return nil, err
	}
	return &entities.GetHealthUserResponse{Status: status}, nil
}

func (s *UserService) Mint(ctx context.Context, address string) error {
	s.log.InfoWithID(ctx, "[Service: Mint] Minting token for address:", address)

	recipient := common.HexToAddress(address)

	// Use the transactor provided via DI
	auth := s.transactor
	auth.Context = ctx

	tx, err := s.userNFTContract.Mint(auth, recipient)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: Mint] Error minting token", err)
		return err
	}
	s.log.InfoWithID(ctx, "[Service: Mint] Token minted successfully", "Transaction Hash", tx.Hash().Hex())
	return nil
}

func (s *UserService) BalanceOf(ctx context.Context, address string) (int64, error) {
	s.log.InfoWithID(ctx, "[Service: BalanceOf] Checking balance for address:", address)
	owner := common.HexToAddress(address)
	balance, err := s.userNFTContract.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, owner) // Read-only call
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: BalanceOf] Error fetching balance", err)
		return 0, err
	}
	return balance.Int64(), nil
}

func (s *UserService) Burn(ctx context.Context, tokenId string) error {
	s.log.InfoWithID(ctx, "[Service: Burn] Burning token with ID:", tokenId)

	tokenIDBigInt := new(big.Int)
	tokenIDBigInt.SetString(tokenId, 10) // Assuming tokenId is provided as a string

	// Use the transactor provided via DI
	auth := s.transactor
	auth.Context = ctx

	tx, err := s.userNFTContract.Burn(auth, tokenIDBigInt)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: Burn] Error burning token", err)
		return err
	}
	s.log.InfoWithID(ctx, "[Service: Burn] Token burned successfully", "Transaction Hash", tx.Hash().Hex())
	return nil
}

func (s *UserService) Login(ctx context.Context, walletAddress string) (*models.User, string, error) {
	s.log.InfoWithID(ctx, "[Service: Login] Logging in user with wallet address:", walletAddress)

	user, err := s.repo.FindByWalletAddress(ctx, walletAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new user with default values
			newUser := &models.User{
				ID:            uuid.New().String(), // Generate a unique ID
				WalletAddress: walletAddress,
				IsActive:      false,
			}
			err = s.repo.CreateUser(ctx, newUser)
			if err != nil {
				s.log.ErrorWithID(ctx, "[Service: Login] Failed to create new user", err)
				return nil, "", err
			}
			user = newUser
		} else {
			s.log.ErrorWithID(ctx, "[Service: Login] Failed to find user by wallet address", err)
			return nil, "", err
		}
	}

	// Generate a JWT token
	tokenString, err := s.generateJWTToken(walletAddress)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: Login] Failed to generate JWT token", err)
		return nil, "", err
	}

	return user, tokenString, nil
}

func (s *UserService) generateJWTToken(walletAddress string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   walletAddress,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.Jwt.JWT_SECRET_KEY))
}

func (s *UserService) Logout(ctx context.Context) error {
	// Perform any cleanup if necessary, such as token invalidation
	s.log.InfoWithID(ctx, "[Service: Logout] User logged out")
	return nil // Return nil or appropriate error if needed
}

func (s *UserService) Register(ctx context.Context, walletAddress, username, firstName, lastName, studentID string) error {
	s.log.InfoWithID(ctx, "[Service: Register] Called with walletAddress:", walletAddress)

	// Find the user by wallet address
	user, err := s.repo.FindByWalletAddress(ctx, walletAddress)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: Register] Failed to find user by wallet address", err)
		return err
	}

	// Update the user with the provided details
	user.Username = username
	user.FirstName = firstName
	user.LastName = lastName
	user.StudentID = studentID
	user.IsActive = true // Mark the account as activated

	// Save the updated user back to the database
	_, err = s.repo.UpdateUserById(ctx, user.ID, user)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: Register] Failed to update user", err)
		return err
	}

	return nil
}

func (s *UserService) IsActivate(ctx context.Context, walletAddress string) (bool, error) {
	s.log.InfoWithID(ctx, "[Service: IsActivate] Called with walletAddress:", walletAddress)

	// Find the user by wallet address
	user, err := s.repo.FindByWalletAddress(ctx, walletAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User does not exist
		}
		s.log.ErrorWithID(ctx, "[Service: IsActivate] Failed to find user by wallet address", err)
		return false, err
	}

	// Check if the user is active
	return user.IsActive, nil
}

func (s *UserService) GetUserInfoByWallet(ctx context.Context, walletAddress string) (entities.GetUserInfoByWalletResponse, error) {
	s.log.InfoWithID(ctx, "[Service: GetUserInfoByWallet] Called with walletAddress:", walletAddress)

	// Find the user by wallet address
	user, err := s.repo.FindByWalletAddress(ctx, walletAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.GetUserInfoByWalletResponse{}, nil // Return an empty response if user does not exist
		}
		s.log.ErrorWithID(ctx, "[Service: GetUserInfoByWallet] Failed to find user by wallet address", err)
		return entities.GetUserInfoByWalletResponse{}, err
	}

	// Map user entity to GetUserInfoByWalletResponse
	response := entities.GetUserInfoByWalletResponse{
		ID:            user.ID,
		WalletAddress: user.WalletAddress,
		Username:      user.Username,
		IsActive:      user.IsActive,
		StudentID:     user.StudentID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		HasVoteYear:   user.HasVoteYear,
		HasVoteRep:    user.HasVoteRep,
	}

	return response, nil
}

func (s *UserService) UpdateUserStatus(ctx context.Context, walletAddress string, req *entities.UpdateUserStatusRequest) error {
	s.log.InfoWithID(ctx, "[Service: UpdateUserStatus] Called with walletAddress:", walletAddress)

	// Build the map of fields to update
	updates := map[string]interface{}{}
	if req.HasVoteYear != nil {
		updates["has_vote_year"] = *req.HasVoteYear
	}
	if req.HasVoteRep != nil {
		updates["has_vote_rep"] = *req.HasVoteRep
	}

	if len(updates) == 0 {
		return nil // Nothing to update
	}

	// Call the repository to apply updates
	err := s.repo.UpdateUser(ctx, walletAddress, updates)
	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: UpdateUserStatus] Failed to update user status", err)
		return err
	}

	return nil
}

func (s *UserService) AdminLogin(ctx context.Context, username, password string) (*models.Admin, error) {
	admin, err := s.repo.FindAdminByUsernameAndPassword(ctx, username, password)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("invalid username or password")
		}
		return nil, err
	}
	return admin, nil
}

func (s *UserService) CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	err := s.repo.CreateCandidate(ctx, candidate)
	if err != nil {
		return nil, err
	}
	return candidate, nil
}

func (s *UserService) GetCandidateByID(ctx context.Context, id string) (*models.Candidate, error) {
	return s.repo.GetCandidateByID(ctx, id)
}

func (s *UserService) UpdateCandidate(ctx context.Context, id string, candidate *models.Candidate) (*models.Candidate, error) {
	return s.repo.UpdateCandidate(ctx, id, candidate)
}

func (s *UserService) DeleteCandidate(ctx context.Context, id string) error {
	return s.repo.DeleteCandidate(ctx, id)
}

func (s *UserService) GetCandidates(ctx context.Context, position string) ([]*models.Candidate, error) {
	var candidates []*models.Candidate
	var err error

	if position != "" {
		// Fetch candidates filtered by position
		candidates, err = s.repo.GetCandidatesByPosition(ctx, position)
	} else {
		// Fetch all candidates
		candidates, err = s.repo.GetAllCandidates(ctx)
	}

	if err != nil {
		s.log.ErrorWithID(ctx, "[Service: GetCandidates] Error fetching candidates", err)
		return nil, err
	}

	return candidates, nil
}
