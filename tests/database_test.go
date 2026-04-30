package tests

import (
	"os"
	"path"
	"testing"
	"wehook-consumer/config"
)

func TestCanConnectDatabase_Success(t *testing.T) {
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()
	if err := database.Connect(); err != nil {
		t.Errorf("Can't connect to database: %v", err)
	}
}

func TestCanCloseDatabaseConnection_Success(t *testing.T) {
	defer func() {
		_, err := os.Open(path.Join(path.Dir("test.db"), "test.db"))

		if err != nil {
			panic(err)
		}

		if err := os.Remove("test.db"); err != nil {
			panic(err)
		}
	}()

	if err := database.Connect(); err != nil {
		t.Errorf("Can't connect to database: %v", err)
	}

	if err := database.Close(); err != nil {
		t.Errorf("Can't close database: %v", err)
	}
}

func TestCanConnectDatabase_Fail(t *testing.T) {

	databaseToFail := config.NewDatabase("test.db", "xpo", false)

	if err := databaseToFail.Connect(); err == nil {
		t.Error("Something went wrong, should throw a error")
	}
}

func init() {
	database = config.NewDatabase("test.db", "sqlite", true)
}
