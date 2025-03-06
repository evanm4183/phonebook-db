package db

import (
	"encoding/json"
	"os"
	"strings"
)

func GetAllContacts() ([]Contact, error) {
	data, err := os.ReadFile("database.txt")
	if err != nil {
		return nil, err
	}

	contacts := []Contact{}
	contactsJson := strings.Split(string(data), "\n")
	contactsJson = contactsJson[:len(contactsJson)-1] // since the last line will always be blank, remove the last record
	for _, jsonData := range contactsJson {
		var c Contact
		json.Unmarshal([]byte(jsonData), &c)
		contacts = append(contacts, c)
	}

	return contacts, nil
}
