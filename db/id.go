package db

import (
	"bufio"
	"os"
	"strconv"
)

func getNextIdThenIncrement() (int, error) {
	file, err := os.OpenFile("id.txt", os.O_RDONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	idText := scanner.Text()
	id, err := strconv.Atoi(idText)
	if err != nil {
		return 0, err
	}

	file, err = os.OpenFile("id.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	_, err = file.WriteString(strconv.Itoa(id + 1))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func resetId() error {
	file, err := os.OpenFile("id.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(1)) // Starting with 1 as first id
	return err
}
