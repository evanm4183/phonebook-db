package db

import (
	"log"
	"os"
)

func WipeDatabase() {
	file, err := os.OpenFile("database.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
