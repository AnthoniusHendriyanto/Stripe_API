package request

type CreateCardDatabaseReq struct {
	CardID     string `json:"cardId"`
	Brand      string `json:"brand"`
	CustomerID string `json:"customerId"`
}

type CreateCardStripeAPI struct {
	CustomerID  string `json:"customerId"`
	PhoneNumber string `json:"phoneNumber"`
	TypeCard    string `json:"typeCard"`
}
type GetListCardsRequest struct {
	Brand       string `json:"brand,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
