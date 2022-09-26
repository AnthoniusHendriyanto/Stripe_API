package stripe

import (
	"fmt"
	"net/http"

	"stripe_api/config"
	"stripe_api/models/request"
	"stripe_api/models/response"
)

type (
	// stripeClientImpl Define a config
	stripeClientImpl struct {
		cfg    *config.Config
		client *http.Client
	}

	Repository interface {
		CreateCustomer(req request.InsertCustomers) (
			*response.CustomerStripeAPIResponse, error,
		)
		CreateCard(req request.CreateCardStripeAPI) (
			*response.CardResponseFromStripe, error,
		)
		Charges(req request.ChargeStripeAPI) (
			*response.ChargesStripeAPIResponse, error,
		)
	}
)

func NewStripeClientImpl() *stripeClientImpl {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
	}

	client := &http.Client{}

	return &stripeClientImpl{
		cfg:    loadConfig,
		client: client,
	}
}
