package ui

import (
	"fmt"
	"phonebook/db"
)

func WipeDatabase() {
	db.WipeDatabase()
	fmt.Println("Database successfully wiped.")
}
