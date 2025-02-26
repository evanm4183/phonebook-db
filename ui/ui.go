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
	actionMap := BuildActionMap()
	actionMap.PrintActions()
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

	action, ok := actionMap[selection]
	if ok {
		action.Invoke()
	} else {
		fmt.Println("Action not found")
	}
}
