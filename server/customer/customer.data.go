package customer

import (
	"database/sql"

	"github.com/didoshotev/XYZ-insurance/database"
)

func getCustomerList() ([]Customer, error) {
	results, err := database.DbConn.Query(`SELECT customerId,
	personName,
	familyName,
	dob,
	rating
	FROM customers`)
	if err != nil {
		return nil, err

	}
	defer results.Close()
	customers := make([]Customer, 0)
	for results.Next() {
		var customer Customer
		results.Scan(&customer.CustomerId,
			&customer.PersonName,
			&customer.FamilyName,
			&customer.DOB,
			&customer.Rating)
		customers = append(customers, customer)
	}
	return customers, nil
}

func getCustomer(customerId int) (*Customer, error) {
	row := database.DbConn.QueryRow(`SELECT customerId,
	personName,
	familyName,
	dob,
	rating
	FROM customers
	WHERE customerId = ?`, customerId)

	customer := &Customer{}
	err := row.Scan(&customer.CustomerId,
		&customer.PersonName,
		&customer.FamilyName,
		&customer.DOB,
		&customer.Rating,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return customer, nil
}

func updateCustomer(customer Customer) (*Customer, error) {
	_, err := database.DbConn.Exec(`Update customers SET
		personName=?,
		familyName=?,
		dob=?,
		rating=?
		WHERE customerId=?`,
		customer.PersonName,
		customer.FamilyName,
		customer.DOB,
		customer.Rating,
		customer.CustomerId)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func createCustomer(customer Customer) (int, error) {
	result, err := database.DbConn.Exec(`INSERT INTO customers
		(personName,
		familyName,
		dob,
		rating) VALUES (?, ?, ?, ?)
		`,
		customer.PersonName,
		customer.FamilyName,
		customer.DOB,
		customer.Rating)
	if err != nil {
		return 0, nil
	}
	createdId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(createdId), nil
}

func removeCustomer(customerId int) error {
	_, err := database.DbConn.Exec(`DELETE FROM customers where customerId = ?`, customerId)
	if err != nil {
		return err
	}
	return nil
}
