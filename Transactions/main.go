package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()

	connStr := "postgres://postgres:1234@localhost:5432/GoBasics"

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatal("not able to start transaction:", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"INSERT INTO orders (user_id, amount) VALUES ($1, $2)",
		1, 3000000,
	)
	if err != nil {
		log.Println("Insert failed:", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Println("Commit failed:", err)
		return
	}

	log.Println("Transaction successful!")
}
