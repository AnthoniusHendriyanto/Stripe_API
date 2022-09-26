package response

type CardResponseFromStripe struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	AddressCity       interface{} `json:"address_city"`
	AddressCountry    interface{} `json:"address_country"`
	AddressLine1      interface{} `json:"address_line1"`
	AddressLine1Check interface{} `json:"address_line1_check"`
	AddressLine2      interface{} `json:"address_line2"`
	AddressState      interface{} `json:"address_state"`
	AddressZip        interface{} `json:"address_zip"`
	AddressZipCheck   interface{} `json:"address_zip_check"`
	Brand             string      `json:"brand"`
	Country           string      `json:"country"`
	Customer          string      `json:"customer"`
	CvcCheck          interface{} `json:"cvc_check"`
	DynamicLast4      interface{} `json:"dynamic_last4"`
	ExpMonth          int         `json:"exp_month"`
	ExpYear           int         `json:"exp_year"`
	Fingerprint       string      `json:"fingerprint"`
	Funding           string      `json:"funding"`
	Last4             string      `json:"last4"`
	Metadata          struct {
	} `json:"metadata"`
	Name               interface{} `json:"name"`
	TokenizationMethod interface{} `json:"tokenization_method"`
}

type CreateCardInternal struct {
	CustomerID  string `json:"customerId"`
	BrandCard   string `json:"brandCard"`
	CardID      string `json:"cardId"`
	ExpiryDate  string `json:"expiryDate"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetListCardInternalResponse struct {
	ListCard []*GetListCards `json:"cards"`
}

type GetListCards struct {
	CardID       string `json:"cardId"`
	BrandCard    string `json:"brandCard"`
	CustomerName string `json:"customerName"`
	PhoneNumber  string `json:"phoneNumber"`
	Status       string `json:"status"`
}
