package ui

import (
	"bufio"
	"os"
	"phonebook/db"
)

func addContact() error {
	reader := bufio.NewReader(os.Stdin)
	contact := db.Contact{}

	contact.FirstName = getInput("First Name: ", reader)
	contact.LastName = getInput("Last Name: ", reader)
	contact.Address = getInput("Address: ", reader)
	contact.PhoneNumber = getInput("Phone Number: ", reader)
	contact.Occupation = getInput("Occupation: ", reader)

	err := db.AddContact(contact)
	if err != nil {
		return err
	}

	return nil
}
