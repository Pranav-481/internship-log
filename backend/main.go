package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"

	"backend/routes"
)

func main() {

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Pranav%401125@localhost:5432/smartcart",
	)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer conn.Close(context.Background())

	routes.SetupRoutes(conn)

	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
