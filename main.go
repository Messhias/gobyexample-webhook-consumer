package main

import "wehook-consumer/config"

func main() {
	database := config.NewDatabase("consumer.db", "sqlite")
	defer func(database config.Database) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)

	if err := database.Connect(); err != nil {
		panic(err)
	}

}
