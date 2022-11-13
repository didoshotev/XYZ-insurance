package main

import (
	"net/http"

	"github.com/didoshotev/XYZ-insurance/auth"
	"github.com/didoshotev/XYZ-insurance/contract"
	"github.com/didoshotev/XYZ-insurance/customer"
	"github.com/didoshotev/XYZ-insurance/database"
	"github.com/didoshotev/XYZ-insurance/insurance"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.SetupDatabase()
	auth.RegisterAuthControllers()
	customer.RegisterCustomerControllers()
	insurance.RegisterInsuranceControllers()
	contract.RegisterContractControllers()
	http.ListenAndServe(":5000", nil)
}
