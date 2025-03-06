package ui

import (
	"bufio"
	"fmt"
	"os"
	"phonebook/db"
	"strconv"
)

func getById() error {
	reader := bufio.NewReader(os.Stdin)

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		return err
	}

	contact, err := db.GetById(id)
	if err != nil {
		return err
	} else if contact != nil {
		fmt.Println(contact)
	} else {
		fmt.Printf("No contact found with Id: %d.", id)
	}

	return nil
}
