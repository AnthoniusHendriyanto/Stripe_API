package constant

const (
	PhoneNumberString   = "phoneNumber"
	BrandString         = "brand"
	PhoneString         = "phone"
	NameString          = "name"
	EmailString         = "email"
	AmountString        = "amount"
	CurrencyString      = "currency"
	CustomerString      = "customer"
	ContentType         = "Content-Type"
	MimeApplicationJSON = "application/json"
	HeaderFormURLEncode = "application/x-www-form-urlencoded"
	HeaderAuthorization = "Authorization"
	BearerTokenString   = "Bearer"
	SourcesString       = "sources"
	SourceString        = "source"
	IDRString           = "IDR"
)

const (
	EndpointCreateCustomer = "v1/customers"
	EndpointCharges        = "v1/charges"
	FormatTwoURL           = "%s/%s"
	FormatFourURL          = "%s/%s/%s/%s"
)

const (
	Inactive = iota
	Active
)
