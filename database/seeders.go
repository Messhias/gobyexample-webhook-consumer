package database

import (
	"gorm.io/gorm"
)

type seeders struct {
	db *gorm.DB
}

func (s seeders) Run() error {
	//TODO implement me
	panic("implement me")
}

type Seeders interface {
	Run() error
}

func NewSeeder(db *gorm.DB) Seeders {
	return &seeders{
		db: db,
	}
}
