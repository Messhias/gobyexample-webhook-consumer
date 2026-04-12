package database

import (
	"wehook-consumer/models"

	"gorm.io/gorm"
)

type Seeders interface {
	Run() error
}

type seeders struct {
	db *gorm.DB
}

func (s seeders) Run() error {
	err := s.seedUsers()

	if err != nil {
		return err
	}

	return nil
}

func NewSeeder(db *gorm.DB) Seeders {
	return &seeders{
		db: db,
	}
}

func (s seeders) seedUsers() error {
	seeds := []models.User{
		{
			Email: "admin@example.com",
		},
		{
			Email: "example@example.com",
		},
	}

	if err := s.db.Create(&seeds).Error; err != nil {
		return err
	}

	return nil
}
