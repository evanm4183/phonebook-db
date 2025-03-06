package ui

import (
	"fmt"
	"phonebook/db"
)

func SeedDatabase() error {
	err := db.SeedDatabase(10)
	if err != nil {
		return err
	} else {
		fmt.Println("Database seeded successfully")
	}

	return nil
}
