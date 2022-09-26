package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"stripe_api/constant"
	"stripe_api/models/request"
)

func (s customerHandler) GetCustomerHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(constant.ContentType, constant.MimeApplicationJSON)

	vars := mux.Vars(req)

	phoneNumber := strings.TrimSpace(vars[constant.PhoneNumberString])

	resp, err := s.services.GetCustomerDetail(request.GetCustomer{
		PhoneNumber: phoneNumber,
	})

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Println("Error Get Customers")
		return
	}

	fmt.Printf("Response: %v\n", resp)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		return
	}
}
