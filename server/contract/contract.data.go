package contract

import (
	"database/sql"

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

func getContract(contractId int) (*Contract, error) {
	row := database.DbConn.QueryRow(`SELECT * FROM contracts WHERE contractId = ?`, contractId)

	contract := &Contract{}
	err := row.Scan(&contract.ContractId,
		&contract.CustomerId,
		&contract.InsuranceId,
		&contract.Details,
		&contract.FinalPrice,
		&contract.Duration,
		&contract.SignUpDate,
		&contract.ValidityDate,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return contract, nil
}

func updateContract(contract Contract) (*Contract, error) {
	_, err := database.DbConn.Exec(`Update contracts SET
		details=?,
		finalPrice=?,
		duration=?,
		validityDate=?
		WHERE contractId=?`,
		contract.Details,
		contract.FinalPrice,
		contract.Duration,
		contract.ValidityDate,
		contract.ContractId)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
