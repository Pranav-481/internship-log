package routes

import (
	"net/http"

	"github.com/jackc/pgx/v5"

	"backend/handlers"
)

func SetupRoutes(conn *pgx.Conn) {

	http.HandleFunc("/api/users", handlers.GetUsers(conn))

	http.HandleFunc("/api/users/create", handlers.CreateUser(conn))
}
