package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {

	connStr := "postgres://postgres:1234@localhost:5432/postgres"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println(err)
	}

	err = conn.Ping(context.Background())

	if err != nil {
		fmt.Println("err", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Connected to PostgreSQL")
}
