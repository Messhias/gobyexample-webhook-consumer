package tests

import (
	"os"
	"path"
	"testing"
	seed "wehook-consumer/database"
	"wehook-consumer/repositories"
)

func TestGetAllUsers_Success(t *testing.T) {
	// arrange
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	err := openDatabase()
	if err != nil {
		t.Errorf("%v", err)
	}

	if database.Seed() {
		seeder := seed.NewSeeder(database.GetDB())
		if err := seeder.Run(); err != nil {
			panic(err)
		}
	}

	repo := repositories.NewUserRepository(database.GetDB())

	// act
	users, err := repo.GetAll()

	// assert
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(users) == 0 {
		t.Error("no users found")
	}

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
