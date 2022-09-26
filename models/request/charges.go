package request

type ChargeStripeAPI struct {
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	CustomerID string `json:"customerId"`
}

type ChargesRequestInternal struct {
	PhoneNumber string `json:"phoneNumber"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
}
