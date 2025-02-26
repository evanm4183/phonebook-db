package ui //test
import (
	"fmt"
	"sort"
)

type Action struct {
	Code   int
	Name   string
	Invoke func()
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

func GetAllActions() [6]Action {
	actions := [6]Action{
		{1, "Get All Records", GetAllRecords},
		{2, "Get Record By ID", GetRecordById},
		{3, "Add Record", AddRecord},
		{4, "Update Record", UpdateRecord},
		{5, "Delete Record", DeleteRecord},
		{6, "Search Records", SearchRecord},
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

func GetAllRecords() {
	fmt.Println("Get All Records")
}

func GetRecordById() {
	fmt.Println("Get Record By ID")
}

func AddRecord() {
	fmt.Println("Add Record")
}

func DeleteRecord() {
	fmt.Println("Delete Record")
}

func UpdateRecord() {
	fmt.Println("Update Record")
}

func SearchRecord() {
	fmt.Println("Search Record")
}
