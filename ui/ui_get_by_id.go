package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"phonebook/db"
	"strconv"
)

func getById() {
	reader := bufio.NewReader(os.Stdin)

	text := getInput("Id: ", reader)
	id, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	contact := db.GetById(id)
	if contact != nil {
		fmt.Println(contact)
	} else {
		fmt.Printf("No contact found with Id: %d.", id)
	}
}
