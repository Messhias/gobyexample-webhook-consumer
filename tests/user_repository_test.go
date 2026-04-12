package tests

import (
	"testing"
	"wehook-consumer/models"
	"wehook-consumer/repositories"
)

func TestGetAllUsers_Success(t *testing.T) {
	// arrange
	err := OpenAndSeed(t, true)

	if err != nil {
		t.Fatal(err)
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

func TestCreateNewUser_Success(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	// arrange
	err := OpenAndSeed(t, false)

	if err != nil {
		t.Fatal(err)
	}

	repo := repositories.NewUserRepository(database.GetDB())

	// arrange
	newUser := models.User{
		Email: "x@x.com",
	}

	// act
	created, err := repo.Create(newUser)

	// assert
	if err != nil {
		t.Error(err)
	}

	if newUser.Email != created.Email {
		t.Error("users doesn't match")
	}
}

func TestCreateNewUserDuplicated_Fail(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	// arrange
	err := OpenAndSeed(t, false)

	if err != nil {
		t.Fatal(err)
	}

	repo := repositories.NewUserRepository(database.GetDB())

	// arrange
	newUser := models.User{
		Email: "admin@example.com",
	}

	// act
	created, err := repo.Create(newUser)

	// assert
	if err != nil {
		t.Error(err)
	}

	if newUser.Email != created.Email {
		t.Error("users doesn't match")
	}

	// act
	_, err = repo.Create(newUser)

	// assert
	if err == nil {
		t.Error("you can't create duplicated user")
	}
}

func TestCreateUserWithWrongEmailShouldFail_Success(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	// arrange
	err := OpenAndSeed(t, false)

	// assert
	if err != nil {
		t.Fatal(err)
	}

	// arrange
	user := models.User{
		Email: "                 ",
	}
	repo := repositories.NewUserRepository(database.GetDB())

	// act
	_, err = repo.Create(user)

	// assert
	if err == nil {
		t.Error("you can't create with wrong email")
	}
}

func TestFindUser_Success(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	// arrange
	err := OpenAndSeed(t, false)

	if err != nil {
		t.Fatal(err)
	}

	repo := repositories.NewUserRepository(database.GetDB())

	// arrange
	newUser := models.User{
		Email: "admin@example.com",
	}

	// act
	created, err := repo.Create(newUser)

	// assert
	if err != nil {
		t.Error(err)
	}

	if newUser.Email != created.Email {
		t.Error("users doesn't match")
	}

	// arrange
	found, err := repo.Find(&newUser)

	if err != nil {
		t.Error(err)
	}

	if found == nil {
		t.Error("user not found")
	}
}

func TestDeleteUser_Success(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	// arrange
	err := OpenAndSeed(t, false)

	if err != nil {
		t.Fatal(err)
	}

	repo := repositories.NewUserRepository(database.GetDB())

	// arrange
	newUser := models.User{
		Email: "admin@example.com",
	}

	// act
	created, err := repo.Create(newUser)

	// assert
	if err != nil {
		t.Error(err)
	}

	if newUser.Email != created.Email {
		t.Error("users doesn't match")
	}

	// arrange
	result, err := repo.Find(&newUser)

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Error("user not found")
		return
	}

	// act
	deleted, err := repo.Delete(result.ExternalID)

	if err != nil {
		t.Error(err)
	}

	if !deleted {
		t.Error("could not delete user")
	}
}
