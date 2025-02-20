package ui

import "fmt"

func Run() {
	fmt.Println("--------------------------")
	fmt.Println("Welcome to the Phone Book!")
	fmt.Println("Select an action:")
	fmt.Println() // would've just added a newline at the end of the preceding statement, but it was throwing a warning and that was annoying.

	actions := GetAllActions()
	for _, a := range actions {
		message := fmt.Sprintf("%v - %v", a.Code, a.Name)
		fmt.Println(message)
	}
	fmt.Println()
}
