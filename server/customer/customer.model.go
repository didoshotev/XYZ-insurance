package models

type Customer struct {
	ID    int
	Name  string
	Type  string
	Price float64
}

var (
	customersCollection []*Customer
	nextCustomerID      = 1
)
