package db

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getNextIdThenIncrement() int {
	file, err := os.OpenFile("id.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	idText := scanner.Text()
	id, err := strconv.Atoi(idText)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.OpenFile("id.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(strconv.Itoa(id + 1))
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func resetId() {
	file, err := os.OpenFile("id.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(strconv.Itoa(1)) // Starting with 1 as first id
}
