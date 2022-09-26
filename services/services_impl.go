package services

import (
	"errors"
	"fmt"

	"stripe_api/constant"
	"stripe_api/models/request"
	"stripe_api/models/response"
)

var errorCustomer = errors.New("STATUS_INACTIVE")

func (s customerService) InsertCustomers(req request.InsertCustomers) (*response.InsertCustomersResponse, error) {
	resp, err := s.InsertCustomerOnStripeAPI(req)

	if err != nil {
		return nil, err
	}

	err = s.SQL.InsertCustomers(request.InsertCustomers{
		CustomerID:  resp.Id,
		Name:        resp.Name,
		PhoneNumber: resp.Phone,
		Status:      req.Status,
	})

	if err != nil {
		fmt.Printf("Error inserting customers : %v", err)
		return nil, err
	}

	status := map[int]string{
		0: "Inactive",
		1: "Active",
	}

	respInternal := &response.InsertCustomersResponse{
		CustomerID:  resp.Id,
		PhoneNumber: resp.Phone,
		Name:        resp.Name,
		Status:      status[req.Status],
		Email:       resp.Email,
	}

	return respInternal, nil
}

func (s customerService) InsertCustomerOnStripeAPI(req request.InsertCustomers) (*response.CustomerStripeAPIResponse, error) {
	resp, err := s.StripeAPI.CreateCustomer(req)
	if err != nil {
		fmt.Printf("Error inserting customers on API Stripe: %v", err)
		return nil, err
	}

	return resp, nil
}

func (s customerService) ChargesStripeAPI(req request.ChargesRequestInternal) (*response.ChargesInternalResponse, error) {
	respCustomer, err := s.GetCustomerDetail(request.GetCustomer{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		fmt.Printf("Error Get customers : %v", err)
		return nil, err
	}

	if respCustomer.Status == constant.Inactive {
		return nil, errorCustomer
	}

	resp, err := s.StripeAPI.Charges(request.ChargeStripeAPI{
		Amount:     req.Amount,
		CustomerID: respCustomer.CustomerID,
	})

	if err != nil {
		fmt.Printf("Error inserting customers on API Stripe: %v", err)
		return nil, err
	}

	respInternal := &response.ChargesInternalResponse{
		Amount:       resp.Amount,
		CustomerID:   resp.Customer,
		ReceiptURL:   resp.ReceiptUrl,
		PhoneNumber:  respCustomer.PhoneNumber,
		CustomerName: respCustomer.Name,
	}

	return respInternal, nil
}

func (s customerService) GetCustomerDetail(req request.GetCustomer) (*response.GetCustomerResponse, error) {
	resp, err := s.SQL.GetCustomers(req)
	if err != nil {
		fmt.Printf("Error Get customers : %v", err)
		return nil, err
	}

	return resp, nil
}

func (s customerService) GetCardDetails(req request.GetListCardsRequest) (*response.GetListCardInternalResponse, error) {
	resp, err := s.SQL.GetCardDetails(req)
	if err != nil {
		fmt.Printf("Error Get Card Details : %v", err)
		return nil, err
	}

	return resp, nil

}

func (s customerService) CreateCardWithExistingCustomer(req request.CreateCardStripeAPI) (*response.CreateCardInternal, error) {
	respCustomer, err := s.GetCustomerDetail(request.GetCustomer{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		fmt.Printf("Error Get customers : %v", err)
		return nil, err
	}

	respCard, err := s.StripeAPI.CreateCard(request.CreateCardStripeAPI{
		CustomerID: respCustomer.CustomerID,
		TypeCard:   req.TypeCard,
	})
	if err != nil {
		fmt.Printf("Error Get customers : %v", err)
		return nil, err
	}

	fmt.Printf("Response Card : %v", respCard)
	expiryDate := fmt.Sprintf("%v-%v", respCard.ExpMonth, respCard.ExpYear)

	err = s.SQL.InsertCard(request.CreateCardDatabaseReq{
		CardID:     respCard.Id,
		Brand:      respCard.Brand,
		CustomerID: respCustomer.CustomerID,
	})

	if err != nil {
		fmt.Printf("Error inserting card on database: %v", err)
		return nil, err
	}

	respInternal := &response.CreateCardInternal{
		CardID:      respCard.Id,
		CustomerID:  respCard.Customer,
		BrandCard:   respCard.Brand,
		ExpiryDate:  expiryDate,
		PhoneNumber: respCustomer.PhoneNumber,
	}

	return respInternal, nil
}
