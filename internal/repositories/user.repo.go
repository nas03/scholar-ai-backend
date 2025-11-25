package repositories

import (
	"context"
	"fmt"

	"github.com/nas03/scholar-ai/backend/internal/consts"
	"github.com/nas03/scholar-ai/backend/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// Basic CRUD operations
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, userID string) (*models.User, error)

	// Update operations
	ActivateUserAccount(ctx context.Context, userID string, status, isEmailVerified int8) error
	UpdateUserAccountStatus(ctx context.Context, userID string, status int8) error
	UpdateUserPassword(ctx context.Context, userID, password string) error
	UpdateUserVerification(ctx context.Context, userID string, isEmailVerified, isPhoneVerified int8) error
	UpdateUser(ctx context.Context, userID string, updates map[string]any) error

	// Transaction operations (like knex.js db.transaction)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	WithTx(tx *gorm.DB) IUserRepository
}

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository with the given database connection.
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// WithTx creates a new instance of the repository with a transaction
func (r *UserRepository) WithTx(tx *gorm.DB) IUserRepository {
	return &UserRepository{db: tx}
}

// CreateUser inserts a new user into the database.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetUserByEmail retrieves a user by email address with optimized query.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Select("user_id, username, email, password, phone_number, account_status, is_email_verified, is_phone_verified, created_at, updated_at").
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by ID with optimized query.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Select("user_id, username, email, password, phone_number, account_status, is_email_verified, is_phone_verified, created_at, updated_at").
		Where("user_id = ?", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user fields.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) UpdateUser(ctx context.Context, userID string, updates map[string]interface{}) error {
	// Remove fields that shouldn't be updated directly
	delete(updates, "user_id")
	delete(updates, "created_at")

	result := r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Updates(updates)

	return result.Error
}

func (r *UserRepository) ActivateUserAccount(ctx context.Context, userID string, status, isEmailVerified int8) error {
	if status != consts.UserAccountStatus.INACTIVE && status != consts.UserAccountStatus.ACTIVE {
		return fmt.Errorf("invalid account status: %d, must be %d (inactive) or %d (active)",
			status, consts.UserAccountStatus.INACTIVE, consts.UserAccountStatus.ACTIVE)
	}

	return r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"account_status":    status,
			"is_email_verified": isEmailVerified,
		}).Error
}

// UpdateUserAccountStatus updates a user's account status.
// Valid status values: consts.UserAccountStatus.INACTIVE (0) or consts.UserAccountStatus.ACTIVE (1)
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) UpdateUserAccountStatus(ctx context.Context, userID string, status int8) error {
	// Validate status value at repository level for data integrity
	if status != consts.UserAccountStatus.INACTIVE && status != consts.UserAccountStatus.ACTIVE {
		return fmt.Errorf("invalid account status: %d, must be %d (inactive) or %d (active)",
			status, consts.UserAccountStatus.INACTIVE, consts.UserAccountStatus.ACTIVE)
	}

	result := r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Update("account_status", status)

	return result.Error
}

// UpdateUserPassword updates a user's password.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) UpdateUserPassword(ctx context.Context, userID, password string) error {
	result := r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Update("password", password)

	return result.Error
}

// UpdateUserVerification updates user verification status.
// Returns raw GORM error - service layer should handle error interpretation
func (r *UserRepository) UpdateUserVerification(ctx context.Context, userID string, isEmailVerified, isPhoneVerified int8) error {
	updates := map[string]interface{}{
		"is_email_verified": isEmailVerified,
		"is_phone_verified": isPhoneVerified,
	}

	result := r.db.WithContext(ctx).Model(&models.User{}).
		Where("user_id = ?", userID).
		Updates(updates)

	return result.Error
}

// WithTransaction executes a function within a database transaction (like knex.js db.transaction)
func (r *UserRepository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
