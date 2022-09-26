package request

type InsertCustomers struct {
	CustomerID  string `json:"customerId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Status      int    `json:"status"`
	Email       string `json:"email,omitempty"`
}

type GetCustomer struct {
	PhoneNumber string `json:"phoneNumber"`
}
