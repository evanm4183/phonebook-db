package ui

import (
	"fmt"
	"phonebook/db"
)

func SeedDatabase() {
	db.SeedDatabase(10)
	fmt.Println("Database seeded successfully")
}
