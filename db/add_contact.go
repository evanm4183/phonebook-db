package db

import (
	"encoding/json"
	"log"
	"os"
)

func AddContact(contact Contact) {
	file, err := os.OpenFile("database.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := json.Marshal(contact)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Writing a newline at the end so that each contact will have it's separate line
	_, err = file.WriteString("\n")
	if err != nil {
		log.Fatal(err)
	}
}
