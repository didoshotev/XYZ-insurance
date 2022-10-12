package models

type Insurance struct {
	ID           int
	CustomerID   int
	InsuranceID  int
	Details      string
	FinalPrice   float64
	Duration     int
	SignUpDate   string
	ValidityDate string
}

var (
	insurancesCollection []*Insurance
	nextInsuranceID      = 1
)
