package stripe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"stripe_api/constant"
	"stripe_api/models/request"
	"stripe_api/models/response"
)

func (s stripeClientImpl) CreateCustomer(req request.InsertCustomers) (
	*response.CustomerStripeAPIResponse, error,
) {
	resp := &response.CustomerStripeAPIResponse{} //nolint:exhaustruct,exhaustivestruct
	createCustomerURL := fmt.Sprintf(constant.FormatTwoURL, s.cfg.StripeAPI.Address, constant.EndpointCreateCustomer)

	client := http.Client{}

	bearerToken := fmt.Sprintf("%v %s", constant.BearerTokenString, s.cfg.StripeAPI.Token)

	urlEncode := url.Values{}

	urlEncode.Set(constant.EmailString, req.Email)
	urlEncode.Set(constant.NameString, req.Name)
	urlEncode.Set(constant.PhoneString, req.PhoneNumber)

	headers := http.Header{
		constant.HeaderAuthorization: []string{bearerToken},
		constant.ContentType:         []string{constant.HeaderFormURLEncode},
	}

	reqStripe, errReq := http.NewRequest(http.MethodPost, createCustomerURL, strings.NewReader(urlEncode.Encode()))

	if errReq != nil {
		fmt.Printf("Error creating customer: %v\n", errReq)
		return nil, errReq
	}

	reqStripe.Header = headers

	fmt.Printf("Request : %v\n", reqStripe)

	res, errResp := client.Do(reqStripe)
	if errResp != nil {
		fmt.Printf("Error from client do : %v\n", errResp)
		return nil, errResp
	}

	resBody, errRead := io.ReadAll(res.Body)

	if errRead != nil {
		fmt.Printf("Error when read body : %v\n", errRead)

		return nil, errRead
	}

	errUnmarshal := json.Unmarshal(resBody, &resp)

	if errUnmarshal != nil {
		fmt.Printf("Error when unmarshalling response : %v\n", errUnmarshal)

		return nil, errUnmarshal
	}

	fmt.Printf("Response : %v\n", resp)

	return resp, nil
}

func (s stripeClientImpl) CreateCard(req request.CreateCardStripeAPI) (
	*response.CardResponseFromStripe, error,
) {
	resp := &response.CardResponseFromStripe{} //nolint:exhaustruct,exhaustivestruct

	// https://api.stripe.com/v1/customers/cus_MT4QAOJ1G4E0Tb/sources
	createCardURL := fmt.Sprintf(constant.FormatFourURL, s.cfg.StripeAPI.Address, constant.EndpointCreateCustomer,
		req.CustomerID, constant.SourcesString)

	fmt.Println("URL: " + createCardURL)

	client := http.Client{}

	bearerToken := fmt.Sprintf("%v %s", constant.BearerTokenString, s.cfg.StripeAPI.Token)

	urlEncode := url.Values{}

	urlEncode.Set(constant.SourceString, req.TypeCard)

	headers := http.Header{
		constant.HeaderAuthorization: []string{bearerToken},
		constant.ContentType:         []string{constant.HeaderFormURLEncode},
	}

	reqStripe, errReq := http.NewRequest(http.MethodPost, createCardURL, strings.NewReader(urlEncode.Encode()))

	if errReq != nil {
		fmt.Printf("Error creating customer: %v\n", errReq)
		return nil, errReq
	}

	reqStripe.Header = headers

	fmt.Printf("Request : %v\n", reqStripe)

	res, errResp := client.Do(reqStripe)
	if errResp != nil {
		fmt.Printf("Error from client do : %v\n", errResp)
		return nil, errResp
	}

	resBody, errRead := io.ReadAll(res.Body)

	if errRead != nil {
		fmt.Printf("Error when read body : %v\n", errRead)

		return nil, errRead
	}

	errUnmarshal := json.Unmarshal(resBody, &resp)

	if errUnmarshal != nil {
		fmt.Printf("Error when unmarshalling response : %v\n", errUnmarshal)

		return nil, errUnmarshal
	}

	fmt.Printf("Response : %v\n", resp)

	return resp, nil
}

func (s stripeClientImpl) Charges(req request.ChargeStripeAPI) (
	*response.ChargesStripeAPIResponse, error,
) {
	resp := &response.ChargesStripeAPIResponse{}
	chargesURL := fmt.Sprintf(constant.FormatTwoURL, s.cfg.StripeAPI.Address, constant.EndpointCharges)

	client := http.Client{}

	bearerToken := fmt.Sprintf("%v %v", constant.BearerTokenString, s.cfg.StripeAPI.Token)

	urlEncode := url.Values{}

	urlEncode.Set(constant.AmountString, req.Amount)
	urlEncode.Set(constant.CurrencyString, constant.IDRString)
	urlEncode.Set(constant.CustomerString, req.CustomerID)

	headers := http.Header{
		constant.HeaderAuthorization: []string{bearerToken},
		constant.ContentType:         []string{constant.HeaderFormURLEncode},
	}

	reqStripe, errReq := http.NewRequest(http.MethodPost, chargesURL, strings.NewReader(urlEncode.Encode()))

	if errReq != nil {
		fmt.Printf("Error when do the payments: %v\n", errReq)
		return nil, errReq
	}

	reqStripe.Header = headers

	fmt.Printf("Request : %v\n", reqStripe)

	res, errResp := client.Do(reqStripe)
	if errResp != nil {
		fmt.Printf("Error from client do : %v\n", errResp)
		return nil, errResp
	}

	resBody, errRead := io.ReadAll(res.Body)

	if errRead != nil {
		fmt.Printf("Error when read body : %v\n", errRead)

		return nil, errRead
	}

	errUnmarshal := json.Unmarshal(resBody, &resp)

	if errUnmarshal != nil {
		fmt.Printf("Error when unmarshalling response : %v\n", errUnmarshal)

		return nil, errUnmarshal
	}

	fmt.Printf("Response : %v\n", resp)

	return resp, nil
}
