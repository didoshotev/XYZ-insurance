package main

import (
	"net/http"

	"github.com/didoshotev/XYZ-insurance/customer"
	"github.com/didoshotev/XYZ-insurance/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.SetupDatabase()
	customer.RegisterCustomerControllers()
	http.ListenAndServe(":5000", nil)
}
