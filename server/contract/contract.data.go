package contract

import (
	"github.com/didoshotev/XYZ-insurance/database"
)

func getContractList() ([]Contract, error) {
	results, err := database.DbConn.Query(`SELECT * FROM contracts`)
	if err != nil {
		return nil, err
	}

	defer results.Close()
	contracts := make([]Contract, 0)
	for results.Next() {
		var contract Contract
		results.Scan(&contract.ContractId,
			&contract.CustomerId,
			&contract.InsuranceId,
			&contract.Details,
			&contract.FinalPrice,
			&contract.Duration,
			&contract.SignUpDate,
			&contract.ValidityDate,
		)
		contracts = append(contracts, contract)
	}
	return contracts, nil
}
