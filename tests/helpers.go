package tests

import (
	"os"
	"path"
	"testing"
	"wehook-consumer/config"
	databaseConfig "wehook-consumer/database"
)

var database config.Database

func OpenAndSeed(t *testing.T, runSeeder bool) error {

	err := openDatabase()
	if err != nil {
		t.Errorf("%v", err)
	}

	if runSeeder {
		if database.Seed() {
			seeder := databaseConfig.NewSeeder(database.GetDB())
			if err := seeder.Run(); err != nil {
				panic(err)
			}
		}
	}

	return err
}

func closeDatabase() error {

	defer func() {
		_, err := os.Open(path.Join(path.Dir("test.db"), "test.db"))

		if err != nil {
			panic(err)
		}

		if err := os.Remove("test.db"); err != nil {
			panic(err)
		}
	}()

	err := database.Close()
	if err != nil {
		return err
	}

	return nil
}

func openDatabase() error {
	err := database.Connect()
	if err != nil {
		return nil
	}
	return err
}
