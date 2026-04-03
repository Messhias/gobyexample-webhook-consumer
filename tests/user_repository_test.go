package tests

import (
	"testing"
	"wehook-consumer/repositories"
)

func TestGetAllUsers_Success(t *testing.T) {
	// arrange
	err := openDatabase()
	if err != nil {
		t.Errorf("%v", err)
	}

	// act
	repo := repositories.NewUserRepository(database.GetDB())
	_, err = repo.GetAll()

	// assert
	if err != nil {
		t.Errorf("%v", err)
	}

	// act
	err = closeDatabase()

	// assert
	if err != nil {
		t.Errorf("%v", err)
	}
}

func closeDatabase() error {
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
