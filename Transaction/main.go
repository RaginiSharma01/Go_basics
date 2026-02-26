package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type transaction struct {
	TransactionID     string    `json:"transactionId"`
	UserId            int       `json:"userId"`
	Amount            float64   `json:"amount"`
	Currency          string    `json:"currency"`
	TransactionStatus string    `json:"transactionStatus"`
	CreatedAt         time.Time `json:"createdAt"`
}

func main() {
	connStr := "postgres://postgres:1234@localhost:5432/GoBasics"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
}
