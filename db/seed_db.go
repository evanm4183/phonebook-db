package db

import (
	"fmt"
	"math/rand"
)

func SeedDatabase(numSeeds int) error {
	contacts := []Contact{}

	for i := 0; i < numSeeds; i++ {
		firstNameIdx := rand.Intn(len(firstNames))
		lastNameIdx := rand.Intn(len(lastNames))
		addressIdx := rand.Intn(len(addresses))
		occupationIdx := rand.Intn(len(occupations))

		contact := Contact{
			i + 1,
			firstNames[firstNameIdx],
			lastNames[lastNameIdx],
			addresses[addressIdx],
			generateRandomPhoneNumber(),
			occupations[occupationIdx],
		}
		contacts = append(contacts, contact)
	}

	for _, c := range contacts {
		err := AddContact(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateRandomPhoneNumber() string {
	// Not exactly how a real phone number would be, but expedient for purposes of this func
	areaCode := rand.Intn(900) + 100
	firstThree := rand.Intn(900) + 100
	lastFour := rand.Intn(9000) + 1000

	return fmt.Sprintf("%d-%d-%d", areaCode, firstThree, lastFour)
}

var firstNames []string = []string{
	"Eddard",
	"Hoster",
	"Jon",
	"Tywin",
	"Robert",
	"Mace",
	"Doran",
}

var lastNames []string = []string{
	"Stark",
	"Tully",
	"Arryn",
	"Lannister",
	"Baratheon",
	"Tyrell",
	"Martell",
}

var addresses []string = []string{
	"Winterfell",
	"Riverrun",
	"The Eyrie",
	"Casterly Rock",
	"Storm's End",
	"High Garden",
	"Sunspear",
}

var occupations []string = []string{
	"King",
	"Lord",
	"Knight",
	"Ranger",
	"Sell Sword",
	"Maester",
	"Septon",
}
