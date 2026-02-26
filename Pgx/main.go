package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Users struct {
	ID    int    `json:"id"`
	User  string `json:"user"`
	Email string `json:"email"`
}

type Post struct {
	ID    int    `json:"id"`
	Posts string `json:"posts"`
	Title string `json:"title"`
}

func main() {
	connStr := "postgres://postgres:1234@localhost:5432/GoBasics"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	names := []string{"vinith", "ragini", "thangasamy"}
	emails := []string{"abc@gmail.com", "abd@gmail.com", "abg@gmail.com"}

	posts := []string{"hello its my first on medium", "Power of young gen :) Students of nepal united and fought for there rights. Fed from corruption!"}
	title := []string{"POV:first post", "Nepal on fireeee!"}

	for i := range names {
		_, err = pool.Exec(context.Background(), "INSERT INTO users (name , email) VALUES ($1, $2) ON CONFLICT DO NOTHING", names[i], emails[i])

		for j := range posts {
			_, err = pool.Exec(context.Background(), "INSERT INTO post (posts, title) VALUES ($1, $2) ON CONFLICT DO NOTHING", posts[j], title[j])

		}
		if err != nil {
			return
		}
	}

	if err != nil {
		log.Fatal("Err in adding the data", err)
	}
}
