package db

import "slices"

func DeleteContact(id int) (bool, error) {
	found := false

	contacts, err := GetAllContacts()
	if err != nil {
		return false, err
	}

	for i := range contacts {
		if contacts[i].Id == id {
			contacts = slices.Delete(contacts, i, i+1)
			found = true
			break
		}
	}
	if !found {
		return false, nil
	}

	err = WipeDatabase()
	if err != nil {
		return false, err
	}

	for _, c := range contacts {
		err = AddContact(c)
		if err != nil {
			return false, err
		}
	}

	return found, nil
}
