package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phonebook/db"
	"strconv"
)

func deleteContact() error {
	reader := bufio.NewReader(os.Stdin)

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	deleteSuccessful, err := db.DeleteContact(id)
	if err != nil {
		return err
	} else if deleteSuccessful {
		fmt.Println("Delete succeeded")
	} else {
		fmt.Printf("Could not find contact with Id: %d\n", id)
	}

	return nil
}
