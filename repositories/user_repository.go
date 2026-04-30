package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"wehook-consumer/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() (*[]models.User, error)
	Create(user models.User) (*models.User, error)
	Delete(id uuid.UUID) (bool, error)
	Find(user *models.User) (*models.User, error)
	Update(id uuid.UUID, user *models.User) (*models.User, error)
}

type userRepository struct {
	database *gorm.DB
	ctx      context.Context
	mu       *sync.Mutex
}

func (u *userRepository) Update(id uuid.UUID, user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if id.String() == "" {
		return nil, errors.New("user external id is empty")
	}

	var foundUser *models.User

	if err := u.database.Where("external_id = ?", id).First(&foundUser).Error; err != nil {
		return nil, err
	}

	foundUser = user

	updatedUser, err := u.Update(user.ExternalID, user)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *userRepository) Find(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	if strings.TrimSpace(user.Email) == "" {
		return nil, errors.New("user's email is empty")
	}

	if strings.TrimSpace(user.ExternalID.String()) == "" {
		return nil, errors.New("user's externalID is empty")
	}

	found, err := gorm.G[models.User](u.database).Where(
		"email = ? or external_id = ?", user.Email, user.ExternalID,
	).First(u.ctx)

	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(found.Email) == "" {
		return nil, errors.New("not valid user provided")
	}

	return &found, nil
}

func (u *userRepository) Delete(id uuid.UUID) (bool, error) {
	if rowsDeleted, err := gorm.G[models.User](u.database).
		Where("external_id = ?", id.String()).
		Delete(u.ctx); err != nil {

		if rowsDeleted == 0 {
			return false, errors.New(fmt.Sprintf("failed to delete user with ID: %s", id.String()))
		}

		return false, err
	}

	return true, nil
}

func (u *userRepository) Create(user models.User) (*models.User, error) {
	if strings.TrimSpace(user.Email) == "" {
		return nil, errors.New("email is required")
	}

	if err := gorm.G[models.User](u.database).Create(u.ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetAll() (*[]models.User, error) {
	users, err := gorm.G[models.User](u.database).Find(u.ctx)

	if err != nil {
		return nil, err
	}

	return &users, nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		database: database,
		ctx:      context.Background(),
	}
}
