package main

import (
	"wehook-consumer/config"
	seed "wehook-consumer/database"
)

func main() {
	database := config.NewDatabase("consumer.db", "sqlite", false)
	defer func(database config.Database) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)

	if err := database.Connect(); err != nil {
		panic(err)
	}

	if database.Seed() {
		seeder := seed.NewSeeder(database.GetDB())
		if err := seeder.Run(); err != nil {
			panic(err)
		}
	}

}
