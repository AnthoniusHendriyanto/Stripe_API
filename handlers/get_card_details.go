package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"stripe_api/constant"
	"stripe_api/models/request"
)

func (s customerHandler) GetCardDetailsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(constant.ContentType, constant.MimeApplicationJSON)

	phoneNumber := req.URL.Query().Get(constant.PhoneNumberString)
	brand := req.URL.Query().Get(constant.BrandString)

	resp, err := s.services.GetCardDetails(request.GetListCardsRequest{
		Brand:       brand,
		PhoneNumber: phoneNumber,
	})

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Error Get Cards")
		return
	}

	fmt.Printf("Response: %v\n", resp)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return
	}
}
