package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetMongoDB()

	r := mux.NewRouter()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc AccountService
	svc = accountservice{}
	{
		repository, err := NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = NewService(repository, logger)
	}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateAccountHandler := httptransport.NewServer(
		makeCreateCustomerEndpoint(svc),
		decodeCreateCustomerRequest,
		encodeResponse,
	)
	GetByIdHandler := httptransport.NewServer(
		makeGetCustomerByIdEndpoint(svc),
		decodeGetCustomerByIdRequest,
		encodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		makeGetAllCustomersEndpoint(svc),
		decodeGetAllCustomersRequest,
		encodeResponse,
	)

	// r.Handle("/hello", )

	// // r.Methods("POST").Path("/account").Handler(CreateAccountHandler)
	// r.Handle("/hello", HelloHandler).Methods("GET")
	http.Handle("/", r)
	http.Handle("/account", CreateAccountHandler)
	//r.Methods("GET").Path("/hello").Handler(HelloHandler)
	r.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	r.Handle("/account/{id}", GetByIdHandler).Methods("GET")

	// http.Handle("/account/:{}", GetByIdHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
