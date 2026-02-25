package main

import (
	"context"
	"fmt"

	"data_access/db"
)

func main() {

	conn := db.ConnectDB()

	defer conn.Close(context.Background())

	db.InsertUser(conn, "Ragini", "ragini@email.com")

	users := db.GetUsers(conn)

	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Email)
	}
}
