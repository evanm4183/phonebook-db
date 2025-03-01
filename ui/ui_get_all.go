package ui

import (
	"fmt"
	"phonebook/db"
)

func getAllContacts() {
	contacts := db.GetAllContacts()
	for _, c := range contacts {
		fmt.Println(c)
	}
}
