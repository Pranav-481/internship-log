package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDatabase() *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Pranav%401125@localhost:5432/smartcart",
	)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Database connected successfully!")

	return conn
}
