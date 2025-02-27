package db

import (
	"fmt"
	"log"
	"os"
)

func AddContact(contact Contact) {
	file, err := os.OpenFile("database.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := fmt.Sprintf("Id: %v\n", contact.Id)
	s += fmt.Sprintf("First Name: %v\n", contact.FirstName)
	s += fmt.Sprintf("Last Name: %v\n", contact.LastName)
	s += fmt.Sprintf("Address: %v\n", contact.Address)
	s += fmt.Sprintf("Phone Number: %v\n", contact.PhoneNumber)
	s += fmt.Sprintf("Occupation: %v\n", contact.Occupation)
	s += "------------\n"

	_, err = file.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
}
