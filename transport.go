package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

// Endpoint for the Account service.

func makeCreateCustomerEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		id, err := s.CreateCustomer(ctx, req.customer)
		return CreateCustomerResponse{Id: id, Err: err}, nil
	}

}

func makeGetCustomerByIdEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetCustomerByIdRequest)
		fmt.Println("in make line 31")
		id, er := strconv.Atoi(req.Id)
		if er != nil {
			return GetCustomerByIdResponse{Email: "", Err: er}, nil
		}
		email, err := s.GetCustomerById(ctx, id)
		return GetCustomerByIdResponse{Email: email, Err: err}, nil
	}

}
func makeGetAllCustomersEndpoint(s AccountService) endpoint.Endpoint {
	fmt.Println("into makeendpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// return GetCustomerByIdResponse{Email: "", Err: nil}, nil
		// req := request.(GetAllCustomersRequest)

		fmt.Println("in make ")
		email, err := s.GetAllCustomers(ctx)
		return GetAllCustomersResponse{Email: email, Err: err}, nil
	}
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req.customer); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCustomerByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetCustomerByIdRequest
	fmt.Println("--------->>>>Into decoding")
	vars := mux.Vars(r)
	req = GetCustomerByIdRequest{
		Id: vars["id"],
	}
	return req, nil
}
func decodeGetAllCustomersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAllCustomersRequest
	fmt.Println("--------->>>>Into GetAllCustomers decoding")
	return req, nil
}

//  encodes the output
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateCustomerRequest struct {
		customer Customer
	}
	CreateCustomerResponse struct {
		Id  string `json:"id"`
		Err error
	}
	GetCustomerByIdRequest struct {
		Id string `json:"id"`
	}
	GetCustomerByIdResponse struct {
		Email string `json:"email"`
		Err   error  `json:"error,omitempty"`
	}
	GetAllCustomersRequest struct {
	}
	GetAllCustomersResponse struct {
		// Id       string `json:"id"`
		Email interface{} `json:"customer"`
		// Phonenum string `json:"phone"`
		Err error `json:"error,omitempty"`
	}
)
