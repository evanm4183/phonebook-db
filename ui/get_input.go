package ui

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal()
	}
	return strings.TrimSpace(text)
}
