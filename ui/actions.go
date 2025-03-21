package ui //test
import (
	"fmt"
	"sort"
)

type Action struct {
	Code   int
	Name   string
	Invoke func() error
}

type ActionMap map[int]Action

func (actionMap ActionMap) PrintActions() {
	if actionMap != nil {
		sorted := []Action{}

		for _, a := range actionMap {
			sorted = append(sorted, a)
		}

		sort.Slice(sorted, func(i int, j int) bool {
			return sorted[i].Code < sorted[j].Code
		})

		for _, a := range sorted {
			fmt.Println(a.Code, "-", a.Name)
		}
	} else {
		fmt.Println("No actions available")
	}
}

func GetAllActions() []Action {
	actions := []Action{
		{1, "Get All Records", getAllContacts},
		{2, "Get Contact By Id", getById},
		{3, "Add Contact", addContact},
		{4, "Update Record", updateContact},
		{5, "Delete Record", deleteContact},
		{6, "Wipe DB", WipeDatabase},
		{7, "Seed DB", SeedDatabase},
	}

	return actions
}

func BuildActionMap() ActionMap {
	m := make(map[int]Action)

	actions := GetAllActions()
	for _, a := range actions {
		m[a.Code] = a
	}

	var actionMap ActionMap = m

	return actionMap
}

func SearchRecord() {
	fmt.Println("Search Record")
}
