package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nnatchy/adv-top-proj/internal/infrastructure/log"
	"github.com/nnatchy/adv-top-proj/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	UpdateLastLogin(context.Context, string) error
	CreateUser(context.Context, *models.User) error
	FindAllUsers(context.Context) (*[]models.User, error)
	UpdateUserById(context.Context, string, *models.User) (*models.User, error)
	DeleteUserById(context.Context, string) error
	FindById(context.Context, string) (*models.User, error)
	FindByStudentId(context.Context, string) (*models.User, error)
	CheckHealth(context.Context) (string, error)
	MintToken(context.Context, string) error
	GetBalance(context.Context, string) (int64, error)
	BurnToken(context.Context, string) error
	FindByWalletAddress(ctx context.Context, walletAddress string) (*models.User, error)
	UpdateUser(ctx context.Context, walletAddress string, updates map[string]interface{}) error
	FindAdminByUsernameAndPassword(ctx context.Context, username, password string) (*models.Admin, error)
	CreateCandidate(ctx context.Context, candidate *models.Candidate) error
	GetCandidateByID(ctx context.Context, id string) (*models.Candidate, error)
	UpdateCandidate(ctx context.Context, id string, candidate *models.Candidate) (*models.Candidate, error)
	DeleteCandidate(ctx context.Context, id string) error
	GetAllCandidates(ctx context.Context) ([]*models.Candidate, error)
	GetCandidatesByPosition(ctx context.Context, position string) ([]*models.Candidate, error)
}

type UserRepository struct {
	db  *gorm.DB
	log *log.Logger
}

func NewUserRepository(
	db *gorm.DB,
	log *log.Logger,
) IUserRepository {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

func (r *UserRepository) FindAllUsers(c context.Context) (*[]models.User, error) {
	r.log.InfoWithID(c, "[Repository: FindAllUsers] Called")
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		r.log.ErrorWithID(c, "[Repository: FindAllUsers] Error fetching users", err)
		return nil, err
	}

	return &users, nil
}

func (r *UserRepository) FindById(c context.Context, id string) (*models.User, error) {
	r.log.InfoWithID(c, "[Repository: FindByID] Called")
	var user models.User

	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		r.log.ErrorWithID(c, "[Repository: FindById] User not found", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByStudentId(c context.Context, id string) (*models.User, error) {
	r.log.InfoWithID(c, "[Repository: FindByID] Called")
	var user models.User

	if err := r.db.Where("student_id = ?", id).First(&user).Error; err != nil {
		r.log.ErrorWithID(c, "[Repository: FindById] User not found", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUserById(c context.Context, id string, updateData *models.User) (*models.User, error) {
	r.log.InfoWithID(c, "[Repository: UpdateUserByID] Called", updateData)

	if result := r.db.Where("id = ?", id).Updates(updateData); result.Error != nil {
		r.log.ErrorWithID(c, "[Repository: UpdateUserByID] Error updating user", result.Error)
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		r.log.ErrorWithID(c, "[Repository: UpdateUserByID] No user updated")
		return nil, errors.New("no user updated")
	}

	return updateData, nil
}

func (r *UserRepository) DeleteUserById(c context.Context, id string) error {
	r.log.InfoWithID(c, "[Repository: DeleteUserById] Called")

	var user models.User

	if err := r.db.Where("user_id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.ErrorWithID(c, "[Repository: DeleteUserById] User not found", err)
			return err
		}
		r.log.ErrorWithID(c, "[Repository: DeleteUserById] Error finding user", err)
		return err
	}

	if err := r.db.Delete(&user).Error; err != nil {
		r.log.ErrorWithID(c, "[Repository: DeleteUserById] Error deleting user", err)
		return err
	}

	r.log.InfoWithID(c, "[Repository: DeleteUserById] User deleted successfully")
	return nil
}

func (r *UserRepository) UpdateLastLogin(c context.Context, email string) error {
	r.log.InfoWithID(c, "[Repository: UpdateLastLogin] Called")
	var user *models.User

	// Use Updates to update specific fields by email
	result := r.db.Model(&user).Where("email = ?", email).Updates(map[string]interface{}{
		"last_login":            time.Now(),
		"failed_login_attempts": 0,
	})

	if result.Error != nil {
		r.log.ErrorWithID(c, "[Repository: UpdateLastLogin] Error updating last login", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.log.ErrorWithID(c, "[Repository: UpdateLastLogin] No rows affected, user not found", result.Error)
		return gorm.ErrRecordNotFound
	}

	return nil
}

//	func (r *UserRepository) CreateUser(c context.Context, user *models.User) error {
//		r.log.InfoWithID(c, "[Repository: CreateUser] Called")
//		if err := r.db.Create(user).Error; err != nil {
//			r.log.ErrorWithID(c, "[Repository: CreateUser] Failed create use in database", err)
//			return errors.New("failed to save user to the database")
//		}
//		return nil
//	}
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) CheckHealth(ctx context.Context) (string, error) {
	// Simulate a health check
	// In real scenarios, this might involve checking database connections, external services, etc.
	return "Healthy", nil
}

func (r *UserRepository) MintToken(ctx context.Context, address string) error {
	// Interaction with smart contract binding for minting
	return nil // Replace with actual logic
}

func (r *UserRepository) GetBalance(ctx context.Context, address string) (int64, error) {
	// Interaction with smart contract binding for balance retrieval
	return 0, nil // Replace with actual logic
}

func (r *UserRepository) BurnToken(ctx context.Context, tokenId string) error {
	// Interaction with smart contract binding for burning
	return nil // Replace with actual logic
}

func (r *UserRepository) FindByWalletAddress(ctx context.Context, walletAddress string) (*models.User, error) {
	var user models.User
	err := r.db.Where("wallet_address = ?", walletAddress).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Generate a unique ID for the new user if ID is a string
			user = models.User{
				ID:            uuid.New().String(), // Ensure ID is set for new user creation
				WalletAddress: walletAddress,
				// Role:          "user", // Default role
				IsActive: false, // Ensure is_active is set to false
			}
			if createErr := r.db.Create(&user).Error; createErr != nil {
				return nil, createErr
			}
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, walletAddress string, updates map[string]interface{}) error {
	r.log.InfoWithID(ctx, "[Repository: UpdateUser] Updating user with wallet address:", walletAddress)

	// Perform a partial update using GORM's Updates()
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(updates).
		Error

	if err != nil {
		r.log.ErrorWithID(ctx, "[Repository: UpdateUser] Error updating user:", err)
		return err
	}

	return nil
}

func (r *UserRepository) FindAdminByUsernameAndPassword(ctx context.Context, username, password string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *UserRepository) CreateCandidate(ctx context.Context, candidate *models.Candidate) error {
	return r.db.WithContext(ctx).Create(candidate).Error
}

func (r *UserRepository) GetCandidateByID(ctx context.Context, id string) (*models.Candidate, error) {
	var candidate models.Candidate
	if err := r.db.WithContext(ctx).First(&candidate, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *UserRepository) UpdateCandidate(ctx context.Context, id string, candidate *models.Candidate) (*models.Candidate, error) {
	if err := r.db.WithContext(ctx).Model(&models.Candidate{}).Where("id = ?", id).Updates(candidate).Error; err != nil {
		return nil, err
	}
	return candidate, nil
}

func (r *UserRepository) DeleteCandidate(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Candidate{}).Error
}

func (r *UserRepository) GetAllCandidates(ctx context.Context) ([]*models.Candidate, error) {
	var candidates []*models.Candidate
	if err := r.db.WithContext(ctx).Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

func (r *UserRepository) GetCandidatesByPosition(ctx context.Context, position string) ([]*models.Candidate, error) {
	var candidates []*models.Candidate
	if err := r.db.WithContext(ctx).Where("position = ?", position).Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}
