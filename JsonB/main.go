package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Settings struct {
	Theme         string `json:"theme"`
	Notifications bool   `json:"notifications"`
}

func main() {
	connStr := "postgres://postgres:1234@localhost:5432/GoBasics"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	arr := []Settings{
		{"dark", true},
		{"blue", true},
		{"light", false},
		{"pink", true},
	}

	for i, s := range arr {
		jsonData, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		_, err = pool.Exec(context.Background(),
			`INSERT INTO user_settings(user_id[i] , settings[i])
	VALUES($1 ,$2)`, i+1, jsonData)
	}
	if err != nil {
		log.Fatal(err)
	}

}
