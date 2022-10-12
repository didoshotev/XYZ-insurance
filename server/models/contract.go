package models

type Contract struct {
	ID         int
	PersonName string
	FamilyName string
	DOB        string
	Rating     int
}

var (
	contractsCollection []*Contract
	nextContractID      = 1
)
