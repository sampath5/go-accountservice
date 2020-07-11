package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           AccountService
}

func (mw instrumentingMiddleware) CreateCustomer(ctx context.Context, customer Customer) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "createcustomer", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.CreateCustomer(ctx, customer)
	return
}

func (mw instrumentingMiddleware) GetCustomerById(ctx context.Context, id int) (Email string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetCustomerById", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	Email, err = mw.next.GetCustomerById(ctx, id)
	return
}

// func (mw instrumentingMiddleware) Count(s string) (n int) {
// 	defer func(begin time.Time) {
// 		lvs := []string{"method", "count", "error", "false"}
// 		mw.requestCount.With(lvs...).Add(1)
// 		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
// 		mw.countResult.Observe(float64(n))
// 	}(time.Now())

// 	n = mw.next.Count(s)
// 	return
// }
