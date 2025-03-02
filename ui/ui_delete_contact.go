package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phonebook/db"
	"strconv"
)

func deleteContact() {
	reader := bufio.NewReader(os.Stdin)

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	deleteSuccessful := db.DeleteContact(id)
	if deleteSuccessful {
		fmt.Println("Delete succeeded")
	} else {
		fmt.Printf("Could not find contact with Id: %d\n", id)
	}
}
