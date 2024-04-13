package main

import "math/rand"

type Account struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        rand.Intn(1000000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(10000000000)),
	}
}
