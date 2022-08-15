package frontend

import (
	"encoding/json"
	"fmt"
	"github.com/quincy0/checkout/component"
	"github.com/quincy0/checkout/util"
)

var BaseURL = "https://api.sandbox.checkout.com/tokens"

type TokenRequest struct {
	Type           string                   `json:"type"`
	Number         string                   `json:"number"`
	ExpiryMonth    int                      `json:"expiry_month"`
	ExpiryYear     int                      `json:"expiry_year"`
	Name           string                   `json:"name,omitempty"`
	Cvv            string                   `json:"cvv,omitempty"`
	BillingAddress component.BillingAddress `json:"billing_address,omitempty"`
	Phone          component.Phone          `json:"phone,omitempty"`
}

type TokenResponse struct {
	Type           string                   `json:"type"`
	Token          string                   `json:"token"`
	ExpiresOn      string                   `json:"expires_on"`
	ExpiryMonth    int                      `json:"expiry_month"`
	ExpiryYear     int                      `json:"expiry_year"`
	Scheme         string                   `json:"scheme,omitempty"`
	SchemeLocal    string                   `json:"scheme_local,omitempty"`
	Last4          string                   `json:"last4"`
	Bin            string                   `json:"bin"`
	CardType       string                   `json:"card_type,omitempty"`
	CardCategory   string                   `json:"card_category,omitempty"`
	Issuer         string                   `json:"issuer,omitempty"`
	IssuerCountry  string                   `json:"issuer_country,omitempty"`
	ProductId      string                   `json:"product_id,omitempty"`
	ProductType    string                   `json:"product_type,omitempty"`
	BillingAddress component.BillingAddress `json:"billing_address,omitempty"`
	Phone          component.Phone          `json:"phone,omitempty"`
	Name           string                   `json:"name,omitempty"`
}

func GetToken() string {
	headers := make(map[string]string)
	headers["Authorization"] = "pk_sbox_t6xr7rbkzd4preonzjcewrilfu="
	headers["Content-Type"] = "application/json"

	parmas := TokenRequest{
		Type:        "card",
		Number:      "4543474002249996",
		ExpiryMonth: 1,
		ExpiryYear:  2038,
		Name:        "Bruce Wayne",
		Cvv:         "956",
		BillingAddress: component.BillingAddress{
			AddressLine1: "Checkout.com",
			AddressLine2: "90 Tottenham Court Road",
			City:         "London",
			State:        "London",
			Zip:          "W1T 4TJ",
			Country:      "GB",
		},
		Phone: component.Phone{CountryCode: "+1", Number: "415 555 2671"},
	}

	jsons, _ := json.Marshal(parmas)
	res, err := util.Post[TokenResponse](BaseURL, jsons, headers)
	//res, err := PostHeader(BaseURL, jsons, headers)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Token)
	return res.Token
}
