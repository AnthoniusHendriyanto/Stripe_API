package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"stripe_api/constant"
	"stripe_api/models/request"
)

func (s customerHandler) CreateCardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constant.ContentType, constant.MimeApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Errors Read All")
		return
	}
	createCard := request.CreateCardStripeAPI{}

	err = json.Unmarshal(body, &createCard)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Println("Errors Unmarshal Create Card")
		return
	}

	fmt.Printf("Request: %v\n", createCard)

	resp, err := s.services.CreateCardWithExistingCustomer(request.CreateCardStripeAPI{
		PhoneNumber: createCard.PhoneNumber,
		TypeCard:    createCard.TypeCard,
	})
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Error inserting card")
		return
	}

	fmt.Printf("Response: %v\n", resp)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return
	}
}
