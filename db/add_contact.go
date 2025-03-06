package db

import (
	"encoding/json"
	"os"
)

func AddContact(contact Contact) error {
	id, err := getNextIdThenIncrement()
	if err != nil {
		return err
	}
	contact.Id = id

	file, err := os.OpenFile("database.txt", os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	// Writing a newline at the end so that each contact will have it's separate line
	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}
