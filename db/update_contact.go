package db

func UpdateContact(contact *Contact) *Contact {
	contacts := GetAllContacts()
	var updated bool = false

	for i := range contacts {
		if contacts[i].Id == contact.Id {
			contacts[i] = *contact
			updated = true
			break
		}
	}
	if !updated {
		return nil
	}

	WipeDatabase()
	for _, c := range contacts {
		AddContact(c)
	}

	return contact
}
