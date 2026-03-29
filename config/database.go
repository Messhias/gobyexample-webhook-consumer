package config

import (
	"errors"
	"fmt"
	"wehook-consumer/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	Connect() error
	Close() error
}

type database struct {
	DB     *gorm.DB
	name   string
	Driver string
}

// Connect implements [Database].
func (d *database) Connect() error {
	connection, err := d.returnConnection()

	if err != nil {
		return err
	}

	db, err := gorm.Open(connection, &gorm.Config{})

	defer func(d *database) {
		err := d.runMigrations()
		if err != nil {
			panic(err)
		}
	}(d)

	if err != nil {
		return err
	}

	d.DB = db

	return nil
}

func (d *database) returnConnection() (gorm.Dialector, error) {
	switch d.Driver {
	case "sqlite":
		return sqlite.Open(d.name), nil

	case "mysql":
	case "sqlserver":
	case "gaussdbgo":
	case "postgres":
	default:
		return nil, errors.New(fmt.Sprintf("%s driver not configured", d.Driver))
	}

	return nil, errors.New(fmt.Sprintf("incorrect driver, %s doesn't exist", d.Driver))
}

func (d *database) Close() error {
	db, err := d.DB.DB()

	if err != nil {
		return err
	}

	return db.Close()
}

func (d *database) runMigrations() error {
	if err := d.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		return err
	}

	return nil
}

func NewDatabase(dbName string, driver string) Database {
	if driver == "" {
		driver = "sqlite"
	}
	return &database{
		name:   dbName,
		Driver: driver,
	}
}
