package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   AccountService
}

func (mw loggingMiddleware) CreateCustomer(ctx context.Context, customer Customer) (Email string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "creatcustomer",
			"Email", customer.Email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Email, err = mw.next.CreateCustomer(ctx, customer)
	return
}

func (mw loggingMiddleware) GetCustomerById(ctx context.Context, id int) (Email string, Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "GetbyId",
			"id", id,
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Email, Err = mw.next.GetCustomerById(ctx, id)
	return
}
func (mw loggingMiddleware) GetAllCustomers(ctx context.Context) (Email interface{}, Err error) {
	defer func(begin time.Time) {

		_ = mw.logger.Log(
			"method", "GetAllCustomers",
			"err", Err,
			"took", time.Since(begin),
		)
	}(time.Now())

	Email, Err = mw.next.GetAllCustomers(ctx)
	return
}
