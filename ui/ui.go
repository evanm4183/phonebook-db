package ui //test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Run() {
	actionMap := BuildActionMap()
	reader := bufio.NewReader(os.Stdin)

	printSeparator()
	fmt.Print("Welcome to the Phone Book\n\n")

	for true {
		actionMap.PrintActions()
		fmt.Println()
		selection := getInput("Select an action: ", reader)
		code, err := strconv.Atoi(selection)
		if err != nil {
			log.Fatal(err)
		}

		action, ok := actionMap[code]
		if ok {
			printSeparator()
			err = action.Invoke()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Action not found")
		}
		printSeparator()
	}
}
