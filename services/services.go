package services

import (
	"stripe_api/models/request"
	"stripe_api/models/response"
	"stripe_api/store/mysql/customers"
	"stripe_api/store/stripe"
)

type customerService struct {
	SQL       customers.Repository
	StripeAPI stripe.Repository
}

type Repository interface {
	InsertCustomers(req request.InsertCustomers) (*response.InsertCustomersResponse, error)
	GetCustomerDetail(req request.GetCustomer) (*response.GetCustomerResponse, error)
	CreateCardWithExistingCustomer(req request.CreateCardStripeAPI) (*response.CreateCardInternal, error)
	GetCardDetails(req request.GetListCardsRequest) (*response.GetListCardInternalResponse, error)
	ChargesStripeAPI(req request.ChargesRequestInternal) (*response.ChargesInternalResponse, error)
}

func NewService(sqlAgent customers.Repository, stripe stripe.Repository) *customerService {
	return &customerService{
		SQL:       sqlAgent,
		StripeAPI: stripe,
	}
}
