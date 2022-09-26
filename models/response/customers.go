package response

type InsertCustomersResponse struct {
	CustomerID  string `json:"customerId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Email       string `json:"email,omitempty"`
}

type GetCustomerResponse struct {
	CustomerID  string `json:"customerId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
}

type CustomerStripeAPIResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Address struct {
		City       string      `json:"city"`
		Country    string      `json:"country"`
		Line1      interface{} `json:"line1"`
		Line2      interface{} `json:"line2"`
		PostalCode string      `json:"postal_code"`
		State      interface{} `json:"state"`
	} `json:"address"`
	Balance         int         `json:"balance"`
	Created         int         `json:"created"`
	Currency        interface{} `json:"currency"`
	DefaultSource   interface{} `json:"default_source"`
	Delinquent      bool        `json:"delinquent"`
	Description     interface{} `json:"description"`
	Discount        interface{} `json:"discount"`
	Email           string      `json:"email"`
	InvoicePrefix   string      `json:"invoice_prefix"`
	InvoiceSettings struct {
		CustomFields         interface{} `json:"custom_fields"`
		DefaultPaymentMethod interface{} `json:"default_payment_method"`
		Footer               interface{} `json:"footer"`
		RenderingOptions     interface{} `json:"rendering_options"`
	} `json:"invoice_settings"`
	Livemode bool `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	Name                string        `json:"name"`
	NextInvoiceSequence int           `json:"next_invoice_sequence"`
	Phone               string        `json:"phone"`
	PreferredLocales    []interface{} `json:"preferred_locales"`
	Shipping            interface{}   `json:"shipping"`
	TaxExempt           string        `json:"tax_exempt"`
	TestClock           interface{}   `json:"test_clock"`
}
