package db

func GetById(id int) (*Contact, error) {
	// TODO: rewrite such that it searches the database and returns only a single record.

	// For simplicity's sake, I'm reusing the GetAllContacts func.
	// Obviously, loading the entire database into memory is not efficient,
	// and there are better ways of implementing this.
	contacts, err := GetAllContacts()
	if err != nil {
		return nil, err
	}

	for _, c := range contacts {
		if c.Id == id {
			return &c, nil
		}
	}

	return nil, nil
}
