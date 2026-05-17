package repository

import (
	"database/sql"
	"log"

	"github.com/bakdaulet212/ap2-final-/user-service/internal/models"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *UserRepository {
	connStr := "host=localhost port=5432 user=user_admin password=my_secret_password dbname=music_app_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (string, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id string
	err := r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}