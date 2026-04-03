package tests

import (
	"testing"
	"wehook-consumer/config"
)

var database config.Database

func TestCanConnectDatabase_Success(t *testing.T) {
	if err := database.Connect(); err != nil {
		t.Errorf("Can't connect to database: %v", err)
	}
}

func TestCanCloseDatabaseConnection_Success(t *testing.T) {

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
	database = config.NewDatabase("test.db", "sqlite", false)
}
