package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"

	"backend/models"
)

func GetUsers(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := conn.Query(
			context.Background(),
			"SELECT id, name, email, password, created_at FROM users",
		)

		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var users []models.User

		for rows.Next() {

			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.CreatedAt,
			)

			if err != nil {
				http.Error(w, "Failed to read user data", http.StatusInternalServerError)
				return
			}

			users = append(users, user)
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(users)
	}
}
func CreateUser(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = conn.QueryRow(
			context.Background(),
			`INSERT INTO users (name, email, password)
			 VALUES ($1, $2, $3)
			 RETURNING id, name, email, password, created_at`,
			user.Name,
			user.Email,
			user.Password,
		).Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)

		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(user)
	}
}
