package ui //test
import "fmt"

type Action struct {
	Code   int
	Name   string
	Invoke func()
}

func GetAllActions() [6]Action {
	actions := [6]Action{
		{1, "Get All Records", GetAllRecords},
		{2, "Get Record By ID", nil},
		{3, "Add Record", nil},
		{4, "Update Record", nil},
		{5, "Delete Record", nil},
		{6, "Search Records", nil},
	}

	return actions
}

func BuildActionsMap() map[int]Action {
	m := make(map[int]Action)

	actions := GetAllActions()
	for _, a := range actions {
		m[a.Code] = a
	}

	return m
}

func GetAllRecords() {
	fmt.Println("All records obtained")
}
