package db

type Contact struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Occupation  string `json:"occupation"`
}
