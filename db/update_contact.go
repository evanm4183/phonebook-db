package db

func UpdateContact(contact *Contact) (*Contact, error) {
	contacts, err := GetAllContacts()
	if err != nil {
		return nil, err
	}
	var updated bool = false

	for i := range contacts {
		if contacts[i].Id == contact.Id {
			contacts[i] = *contact
			updated = true
			break
		}
	}
	if !updated {
		return nil, nil
	}

	err = WipeDatabase()
	if err != nil {
		return nil, err
	}
	for _, c := range contacts {
		err = AddContact(c)
		if err != nil {
			return nil, err
		}
	}

	return contact, nil
}
