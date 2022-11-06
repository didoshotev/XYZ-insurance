package contract

import "time"

type Contract struct {
	ContractId   int       `json:"contractId"`
	CustomerId   int       `json:"customerId"`
	InsuranceId  int       `json:"insuranceId"`
	Details      string    `json:"details"`
	FinalPrice   float64   `json:"finalPrice"`
	Duration     int       `json:"duration"`
	SignUpDate   time.Time `json:"signUpDate"`
	ValidityDate time.Time `json:"validityDate"`
}
