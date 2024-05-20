package main

import (
	"math/rand"
	"time"
)

type CreateAccountReq struct {
	FirstName string 	`json:"firstname"`
	LastName string 	`json:"lastname"`
}

type Account struct {
	ID        int `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Balance   float64 `json:"balance"`
	Number    int `json:"number"`
	CreatedAt time.Time `json:"createdat"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		//ID:        rand.Intn(10000),
		Number:    rand.Intn(10000),
		CreatedAt: time.Now().UTC(),
	}
}
