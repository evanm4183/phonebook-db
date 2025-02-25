package ui //test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("--------------------------")
	fmt.Println("Welcome to the Phone Book")
	fmt.Println("Select an action:")
	fmt.Println() // would've just added a newline at the end of the preceding statement, but it was throwing a warning and that was annoying.
	actions := GetAllActions()
	for _, a := range actions {
		message := fmt.Sprintf("%v - %v", a.Code, a.Name)
		fmt.Println(message)
	}
	fmt.Println()

	fmt.Print("Enter your action: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	trimmed := strings.TrimSpace(text)
	selection, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatal(err)
	}

	switch selection {
	case 1:
		fmt.Println("Case 1")
		actions[0].Invoke()
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	case 4:
		fmt.Println("Case 4")
	case 5:
		fmt.Println("Case 5")
	case 6:
		fmt.Println("Case 6")
	default:
		fmt.Println("Option does not exist")
	}
}
