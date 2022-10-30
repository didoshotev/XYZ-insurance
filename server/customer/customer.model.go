package customer

type Customer struct {
	CustomerId int    `json:"customerId"`
	PersonName string `json:"personName"`
	FamilyName string `json:"familyName"`
	DOB        string `json:"dob"`
	Rating     int    `json:"rating"`
}
