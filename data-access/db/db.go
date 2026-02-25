package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func InsertUser(conn *pgx.Conn, name string, email string) {
	query := `
	INSERT INTO users (name , email)
	VALUES ($1 , $2)`

	_, err := conn.Exec(context.Background(), query, name, email)

	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	log.Println("User inserted")
}

func GetUsers(conn *pgx.Conn) []User {

	query := `SELECT id, name, email FROM users`

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []User

	for rows.Next() {

		var u User

		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
		)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	return users
}

func ConnectDB() *pgx.Conn {

	connStr := "postgres://postgres:1234@localhost:5432/GoBasics"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	fmt.Println("DB Connected from package")

	return conn
}
