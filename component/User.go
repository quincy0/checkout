package component

type BillingAddress struct {
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Country      string `json:"country"`
}

type Phone struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
