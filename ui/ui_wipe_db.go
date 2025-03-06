package ui

import (
	"fmt"
	"phonebook/db"
)

func WipeDatabase() error {
	err := db.WipeDatabase()
	if err != nil {
		return err
	} else {
		fmt.Println("Database successfully wiped.")
	}

	return nil
}
