package response

type ChargesStripeAPIResponse struct {
	Id                   string      `json:"id"`
	Object               string      `json:"object"`
	Amount               int         `json:"amount"`
	AmountCaptured       int         `json:"amount_captured"`
	AmountRefunded       int         `json:"amount_refunded"`
	Application          interface{} `json:"application"`
	ApplicationFee       interface{} `json:"application_fee"`
	ApplicationFeeAmount interface{} `json:"application_fee_amount"`
	BalanceTransaction   string      `json:"balance_transaction"`
	BillingDetails       struct {
		Address struct {
			City       interface{} `json:"city"`
			Country    interface{} `json:"country"`
			Line1      interface{} `json:"line1"`
			Line2      interface{} `json:"line2"`
			PostalCode interface{} `json:"postal_code"`
			State      interface{} `json:"state"`
		} `json:"address"`
		Email interface{} `json:"email"`
		Name  interface{} `json:"name"`
		Phone interface{} `json:"phone"`
	} `json:"billing_details"`
	CalculatedStatementDescriptor string      `json:"calculated_statement_descriptor"`
	Captured                      bool        `json:"captured"`
	Created                       int         `json:"created"`
	Currency                      string      `json:"currency"`
	Customer                      string      `json:"customer"`
	Description                   string      `json:"description"`
	Destination                   interface{} `json:"destination"`
	Dispute                       interface{} `json:"dispute"`
	Disputed                      bool        `json:"disputed"`
	FailureBalanceTransaction     interface{} `json:"failure_balance_transaction"`
	FailureCode                   interface{} `json:"failure_code"`
	FailureMessage                interface{} `json:"failure_message"`
	FraudDetails                  struct {
	} `json:"fraud_details"`
	Invoice  interface{} `json:"invoice"`
	Livemode bool        `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	OnBehalfOf interface{} `json:"on_behalf_of"`
	Order      interface{} `json:"order"`
	Outcome    struct {
		NetworkStatus string      `json:"network_status"`
		Reason        interface{} `json:"reason"`
		RiskLevel     string      `json:"risk_level"`
		RiskScore     int         `json:"risk_score"`
		SellerMessage string      `json:"seller_message"`
		Type          string      `json:"type"`
	} `json:"outcome"`
	Paid                 bool        `json:"paid"`
	PaymentIntent        interface{} `json:"payment_intent"`
	PaymentMethod        string      `json:"payment_method"`
	PaymentMethodDetails struct {
		Card struct {
			Brand  string `json:"brand"`
			Checks struct {
				AddressLine1Check      interface{} `json:"address_line1_check"`
				AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
				CvcCheck               interface{} `json:"cvc_check"`
			} `json:"checks"`
			Country      string      `json:"country"`
			ExpMonth     int         `json:"exp_month"`
			ExpYear      int         `json:"exp_year"`
			Fingerprint  string      `json:"fingerprint"`
			Funding      string      `json:"funding"`
			Installments interface{} `json:"installments"`
			Last4        string      `json:"last4"`
			Mandate      interface{} `json:"mandate"`
			Network      string      `json:"network"`
			ThreeDSecure interface{} `json:"three_d_secure"`
			Wallet       interface{} `json:"wallet"`
		} `json:"card"`
		Type string `json:"type"`
	} `json:"payment_method_details"`
	ReceiptEmail  string      `json:"receipt_email"`
	ReceiptNumber interface{} `json:"receipt_number"`
	ReceiptUrl    string      `json:"receipt_url"`
	Refunded      bool        `json:"refunded"`
	Refunds       struct {
		Object     string        `json:"object"`
		Data       []interface{} `json:"data"`
		HasMore    bool          `json:"has_more"`
		TotalCount int           `json:"total_count"`
		Url        string        `json:"url"`
	} `json:"refunds"`
	Review   interface{} `json:"review"`
	Shipping interface{} `json:"shipping"`
	Source   struct {
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
	} `json:"source"`
	SourceTransfer            interface{} `json:"source_transfer"`
	StatementDescriptor       interface{} `json:"statement_descriptor"`
	StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	Status                    string      `json:"status"`
	TransferData              interface{} `json:"transfer_data"`
	TransferGroup             interface{} `json:"transfer_group"`
}

type ChargesInternalResponse struct {
	Amount       int    `json:"amount"`
	CustomerID   string `json:"customerId"`
	ReceiptURL   string `json:"receiptURL"`
	PhoneNumber  string `json:"phoneNumber"`
	CustomerName string `json:"customerName"`
}
