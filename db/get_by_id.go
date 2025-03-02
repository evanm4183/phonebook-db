package db

func GetById(id int) *Contact {
	// TODO: rewrite such that it searches the database and returns only a single record.

	// For simplicity's sake, I'm reusing the GetAllContacts func.
	// Obviously, loading the entire database into memory is not efficient,
	// and there are better ways of implementing this.
	contacts := GetAllContacts()

	for _, c := range contacts {
		if c.Id == id {
			return &c
		}
	}

	return nil
}
