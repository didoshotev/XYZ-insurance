package insurance

type Insurance struct {
	InsuranceId int     `json:"insuranceId"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
}
