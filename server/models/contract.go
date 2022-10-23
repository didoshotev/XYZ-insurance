package models

import (
	"errors"
	"fmt"
)

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

func GetContractByID(id int) (Contract, error) {
	for _, currContract := range contractsCollection {
		if currContract.ID == id {
			return *currContract, nil
		}
	}
	return Contract{}, fmt.Errorf("contract with such id '%v' does not exists", id)
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

func UpdateContract(contract Contract) (Contract, error) {
	for i, currContract := range contractsCollection {
		if contract.ID == currContract.ID {
			contractsCollection[i] = &contract
			return contract, nil
		}
	}
	return Contract{}, fmt.Errorf("contract not found")
}

func DeleteContract(id int) error {
	for i, currContract := range contractsCollection {
		if currContract.ID == id {
			contractsCollection = append(contractsCollection[:i], contractsCollection[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with such ID '%v' does not exists", id)
}
