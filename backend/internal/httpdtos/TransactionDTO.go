package httpdtos

import "github.com/google/uuid"

type TransactionDTO struct {
	TransactionId uuid.UUID `json:"transactionId"`
	From          string    `json:"from"`
	SessionJwt    string    `json:"sessionJwt"`
	To            string    `json:"to"`
	Value         string    `json:"value"`
	Currency      string    `json:"currency"`
}
