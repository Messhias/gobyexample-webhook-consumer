package tests

import (
	"testing"
	db "wehook-consumer/database"
)

func TestItSeederRun_Success(t *testing.T) {
	// arrange
	err := openDatabase()
	defer func() {
		err := closeDatabase()
		if err != nil {
			t.Errorf("%v", err)
		}
	}()
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	seeder := db.NewSeeder(database.GetDB())

	// act
	err = seeder.Run()

	// assert
	if err != nil {
		t.Fatalf("failed to run seeder: %v", err)
	}
}
