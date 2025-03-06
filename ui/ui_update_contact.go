package ui

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/db"
	"strconv"
)

func updateContact() error {
	reader := bufio.NewReader(os.Stdin)
	contact := db.Contact{}

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		return err
	}
	contact.Id = id
	contact.FirstName = getInput("First Name: ", reader)
	contact.LastName = getInput("Last Name: ", reader)
	contact.Address = getInput("Address: ", reader)
	contact.PhoneNumber = getInput("Phone Number: ", reader)
	contact.Occupation = getInput("Occupation: ", reader)

	result, err := db.UpdateContact(&contact)
	if err != nil {
		return err
	} else if result == nil {
		fmt.Printf("Contact not found with Id: %d\n", contact.Id)
	} else {
		fmt.Println(*result)
	}

	return nil
}
