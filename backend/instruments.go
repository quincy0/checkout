package backend

import (
	"encoding/json"
	"fmt"
	"github.com/quincy0/checkout/component"
	"github.com/quincy0/checkout/util"
	"log"
)

type CreateInstrumentRequest struct {
	Type          string         `json:"type"`
	Token         string         `json:"token"`
	AccountHolder *AccountHolder `json:"account_holder,omitempty"`
	Customer      *Customer      `json:"customer,omitempty"`
}

type AccountHolder struct {
	FirstName      string                   `json:"first_name,omitempty"`
	LastName       string                   `json:"last_name,omitempty"`
	BillingAddress component.BillingAddress `json:"billing_address,omitempty"`
	Phone          component.Phone          `json:"phone,omitempty"`
}

type Customer struct {
	Id      string           `json:"id,omitempty"`
	Email   string           `json:"email,omitempty"`
	Name    string           `json:"name,omitempty"`
	Phone   *component.Phone `json:"phone,omitempty"`
	Default bool             `json:"default,omitempty"`
}

type CreateInstrumentResponse struct {
	Type          string   `json:"type"`
	Id            string   `json:"id"`
	Fingerprint   string   `json:"fingerprint"`
	ExpiryMonth   int      `json:"expiry_month"`
	ExpiryYear    int      `json:"expiry_year"`
	Last4         string   `json:"last4"`
	Bin           string   `json:"bin"`
	Scheme        string   `json:"scheme,omitempty"`
	SchemeLocal   string   `json:"scheme_local,omitempty"`
	CardType      string   `json:"card_type,omitempty"`
	CardCategory  string   `json:"card_category"`
	Issuer        string   `json:"issuer,omitempty"`
	IssuerCountry string   `json:"issuer_country,omitempty"`
	ProductId     string   `json:"product_id,omitempty"`
	ProductType   string   `json:"product_type,omitempty"`
	Customer      Customer `json:"customer,omitempty"`
}

func CreateInstrument(token string, email string) {
	url := "https://api.sandbox.checkout.com/instruments"
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer sk_sbox_ihewjgjywrefafvqcjkex2xz4aa"
	headers["Content-Type"] = "application/json"

	params := CreateInstrumentRequest{
		Type:          "token",
		Token:         token,
		AccountHolder: nil,
		Customer:      &Customer{Email: email},
	}

	data, _ := json.Marshal(params)
	fmt.Println(string(data))
	res, err := util.Post[CreateInstrumentResponse](url, data, headers)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", res)
	log.Println(res)
}
