package ui

type Action struct {
	Code int
	Name string
}

func GetAllActions() [6]Action {
	actions := [6]Action{
		{1, "Get All Records"},
		{2, "Get Record By ID"},
		{3, "Add Record"},
		{4, "Update Record"},
		{5, "Delete Record"},
		{6, "Search Records"},
	}

	return actions
}
