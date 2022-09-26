package customers

import (
	"database/sql"
	"fmt"

	"stripe_api/models/request"
	"stripe_api/models/response"
	"stripe_api/pkg/database"
)

type customerClient struct {
	DB *sql.DB
}

type Repository interface {
	InsertCustomers(req request.InsertCustomers) error
	GetCustomers(req request.GetCustomer) (
		*response.GetCustomerResponse, error)
	InsertCard(req request.CreateCardDatabaseReq) error
	GetCardDetails(req request.GetListCardsRequest) (
		*response.GetListCardInternalResponse, error)
}

func NewClient() *customerClient {
	conn, err := database.DBConnect()

	if err != nil {
		fmt.Printf("Error creating database connection: %v\n", err)
		return nil
	}

	return &customerClient{
		DB: conn,
	}
}
