package repositories

import (
	"context"
	"wehook-consumer/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
}

type userRepository struct {
	database *gorm.DB
	ctx      context.Context
}

func (u userRepository) GetAll() ([]models.User, error) {
	users, err := gorm.G[models.User](u.database).Find(u.ctx)

	if err != nil {
		return []models.User{}, nil
	}

	return users, nil
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepository{
		database: database,
		ctx:      context.Background(),
	}
}
