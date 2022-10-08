package insurance

import (
	"database/sql"
	"log"

	"github.com/didoshotev/XYZ-insurance/database"
)

func getInsuranceList() ([]Insurance, error) {
	results, err := database.DbConn.Query(`SELECT insuranceId,
	name,
	type,
	price
	FROM insurances`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer results.Close()
	insurances := make([]Insurance, 0)
	for results.Next() {
		var insurance Insurance
		results.Scan(&insurance.InsuranceId,
			&insurance.Name,
			&insurance.Type,
			&insurance.Price)
		insurances = append(insurances, insurance)
	}
	return insurances, nil
}

func getInsurance(insuranceId int) (*Insurance, error) {
	row := database.DbConn.QueryRow(`SELECT insuranceId,
	name,
	type,
	price
	FROM insurances
	WHERE insuranceId = ?`, insuranceId)

	insurance := &Insurance{}
	err := row.Scan(&insurance.InsuranceId,
		&insurance.Name,
		&insurance.Type,
		&insurance.Price,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return insurance, nil
}

func updateInsurance(insurance Insurance) (*Insurance, error) {
	_, err := database.DbConn.Exec(`Update insurances SET
		name=?,
		type=?,
		price=?
		WHERE insuranceId=?`,
		&insurance.Name,
		&insurance.Type,
		&insurance.Price,
		insurance.InsuranceId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &insurance, nil
}

func createInsurance(insurance Insurance) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO insurances
		(name,
		type,
		price) VALUES (?, ?, ?)`,
		&insurance.Name,
		&insurance.Type,
		&insurance.Price)
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	createdId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, nil
	}
	return int(createdId), nil
}

func removeInsurance(insuranceId int) error {
	_, err := database.DbConn.Exec(`DELETE FROM insurances where insuranceId = ?`, insuranceId)
	if err != nil {
		return err
	}
	return nil
}
