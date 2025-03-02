package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phonebook/db"
	"strconv"
)

func updateContact() {
	reader := bufio.NewReader(os.Stdin)
	contact := db.Contact{}

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal()
	}
	contact.Id = id
	contact.FirstName = getInput("First Name: ", reader)
	contact.LastName = getInput("Last Name: ", reader)
	contact.Address = getInput("Address: ", reader)
	contact.PhoneNumber = getInput("Phone Number: ", reader)
	contact.Occupation = getInput("Occupation: ", reader)

	result := db.UpdateContact(&contact)
	if result == nil {
		fmt.Printf("Contact not found with Id: %d\n", contact.Id)
	} else {
		fmt.Println(*result)
	}
}
