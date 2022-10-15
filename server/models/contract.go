package models

import "errors"

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

func GetContracts() []*Contract {
	return contractsCollection
}

func CreateContract(newContract Contract) (Contract, error) {
	if newContract.ID != 0 {
		return Contract{}, errors.New("new user must not include ID field or must be set to 0")
	}
	newContract.ID = nextContractID
	nextContractID++
	contractsCollection = append(contractsCollection, &newContract)
	return newContract, nil
}
