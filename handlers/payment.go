package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"stripe_api/constant"
	"stripe_api/models/request"
)

func (s customerHandler) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constant.ContentType, constant.MimeApplicationJSON)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Errors Read All")
		return
	}
	charge := request.ChargesRequestInternal{}

	err = json.Unmarshal(body, &charge)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		fmt.Println("Errors Unmarshal Create Card")
		return
	}

	fmt.Printf("Request: %v\n", charge)

	resp, err := s.services.ChargesStripeAPI(request.ChargesRequestInternal{
		PhoneNumber: charge.PhoneNumber,
		Amount:      charge.Amount,
	})

	fmt.Printf("Err: %v\n", err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Errors From Services")

		return
	}

	fmt.Printf("Response: %v\n", resp)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return
	}
}
