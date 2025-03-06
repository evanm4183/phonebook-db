package ui

import (
	"fmt"
	"phonebook/db"
)

func getAllContacts() error {
	contacts, err := db.GetAllContacts()
	if err != nil {
		return err
	}

	for _, c := range contacts {
		fmt.Println(c)
	}

	return nil
}
