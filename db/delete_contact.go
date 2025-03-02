package db

import "slices"

func DeleteContact(id int) bool {
	found := false

	contacts := GetAllContacts()
	for i := range contacts {
		if contacts[i].Id == id {
			contacts = slices.Delete(contacts, i, i+1)
			found = true
			break
		}
	}
	if !found {
		return false
	}
	WipeDatabase()
	for _, c := range contacts {
		AddContact(c)
	}

	return found
}
