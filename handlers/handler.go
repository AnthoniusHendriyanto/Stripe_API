package handlers

import (
	"net/http"

	"stripe_api/services"
)

type customerHandler struct {
	services services.Repository
}

type Repository interface {
	CreateCustomerHandler(w http.ResponseWriter, r *http.Request)
	GetCustomerHandler(w http.ResponseWriter, r *http.Request)
	CreateCardHandler(w http.ResponseWriter, r *http.Request)
	GetCardDetailsHandler(w http.ResponseWriter, req *http.Request)
	PaymentHandler(w http.ResponseWriter, r *http.Request)
}

func New(s services.Repository) *customerHandler {
	return &customerHandler{
		services: s,
	}
}
