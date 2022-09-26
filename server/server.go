package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"stripe_api/handlers"
	"stripe_api/services"
	"stripe_api/store/mysql/customers"
	"stripe_api/store/stripe"
)

type Server struct {
	services services.Repository
	handlers handlers.Repository
	sql      customers.Repository
	stripe   stripe.Repository
}

func Register() *Server {
	SVR := &Server{}

	SVR.sql = customers.NewClient()

	SVR.stripe = stripe.NewStripeClientImpl()

	SVR.services = services.NewService(SVR.sql, SVR.stripe)

	SVR.handlers = handlers.New(SVR.services)

	return SVR
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/customer", s.handlers.CreateCustomerHandler).Methods(http.MethodPost)
	r.HandleFunc("/card", s.handlers.CreateCardHandler).Methods(http.MethodPost)
	r.HandleFunc("/charges", s.handlers.PaymentHandler).Methods(http.MethodPost)

	r.HandleFunc("/customer/{phoneNumber}", s.handlers.GetCustomerHandler).Methods(http.MethodGet)
	r.HandleFunc("/cards", s.handlers.GetCardDetailsHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error ListenAndServe")
		return
	}
}
