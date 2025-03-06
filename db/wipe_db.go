package db

import (
	"os"
)

func WipeDatabase() error {
	file, err := os.OpenFile("database.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = resetId()
	return err
}
