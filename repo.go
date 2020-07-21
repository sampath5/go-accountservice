package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var RepoErr = errors.New("Unable to handle Repo Request")

const UserCollection = "User"

type repo struct {
	db     *mgo.Database
	logger log.Logger
}

func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

func (repo *repo) CreateCustomer(ctx context.Context, customer Customer) error {
	err := db.C(UserCollection).Insert(customer)
	if err != nil {
		fmt.Println("Error occured inside CreateCustomer in repo")
		return err
	} else {
		fmt.Println("User Created:", customer.Email)
	}
	return nil
}

func (repo *repo) GetCustomerById(ctx context.Context, id int) (string, error) {
	coll := db.C(UserCollection)
	email := Customer{}
	err := coll.Find(bson.M{"customerid": id}).Select(bson.M{"email": 1}).One(&email)
	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return email.Email, nil
}
func (erpo *repo) GetAllCustomers(ctx context.Context) (interface{}, error) {
	coll := db.C(UserCollection)
	email := []Customer{}
	// fmt.Println("into Get AllCustomers repo code")
	err := coll.Find(bson.M{}).Select(bson.M{"id": 1, "customerid": 1, "email": 1, "phone": 1}).All(&email)

	if err != nil {
		fmt.Println("Error occured inside GetCUstomerById in repo")
		return "", err
	}
	return email, nil
}
