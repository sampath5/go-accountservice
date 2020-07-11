package main

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type Customer struct {
	Id         string `json:"id"`
	Customerid int    `json:"customerid"`
	Email      string ` json:"email"`
	Password   string ` json:"password"`
	Phone      string ` json:"phone"`
}

type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	GetCustomerById(ctx context.Context, id int) (string, error)
}

// service implements the ACcount Service
type accountservice struct {
	repository Repository
	logger     log.Logger
}

// Service describes the Account service.
type AccountService interface {
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	GetCustomerById(ctx context.Context, id int) (string, error)
}

// NewService creates and returns a new Account service instance
func NewService(rep Repository, logger log.Logger) AccountService {
	return &accountservice{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an customer
func (s accountservice) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	var id = uuid.String()
	customer.Id = id
	customerDetails := Customer{
		Id:         customer.Id,
		Customerid: customer.Customerid,
		Email:      customer.Email,
		Password:   customer.Password,
		Phone:      customer.Phone,
	}
	println("in BL")
	if err := s.repository.CreateCustomer(ctx, customerDetails); err != nil {
		level.Error(logger).Log("err", err)
	}
	return id, nil
}

func (s accountservice) GetCustomerById(ctx context.Context, id int) (string, error) {
	logger := log.With(s.logger, "method", "GetCustomerById")
	var email string
	email, err := s.repository.GetCustomerById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return email, nil
}
