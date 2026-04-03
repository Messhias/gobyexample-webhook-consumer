package tests

import (
	"testing"
	"wehook-consumer/repositories"
)

func TestGetAllUsers_Success(t *testing.T) {
	// arrange
	err, _ := openDatabase()
	if err != nil {
		t.Errorf("%v", err)
	}

	repo := repositories.NewUserRepository(database.GetDB())
	_, err = repo.GetAll()

	if err != nil {
		t.Errorf("%v", err)
	}

	err = closeDatabase()

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

func openDatabase() (error, bool) {
	err := database.Connect()
	if err != nil {
		return nil, true
	}
	return err, false
}
