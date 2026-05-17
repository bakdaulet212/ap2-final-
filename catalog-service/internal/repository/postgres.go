package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/bakdaulet212/ap2-final-/catalog-service/internal/models"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type CatalogRepository struct {
	db  *sql.DB
	rdb *redis.Client
	ctx context.Context
}

func NewCatalogRepository() *CatalogRepository {
	connStr := "host=localhost port=5432 user=user_admin password=my_secret_password dbname=music_app_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})

	query := `
	CREATE TABLE IF NOT EXISTS tracks (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(100) NOT NULL,
		artist VARCHAR(100) NOT NULL,
		duration VARCHAR(10) NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tracks table: %v", err)
	}

	return &CatalogRepository{
		db:  db,
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *CatalogRepository) AddTrack(track *models.Track) (string, error) {
	query := `INSERT INTO tracks (title, artist, duration) VALUES ($1, $2, $3) RETURNING id`
	var id string
	err := r.db.QueryRow(query, track.Title, track.Artist, track.Duration).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *CatalogRepository) GetTrack(id string) (*models.Track, error) {
	cachedTrack, err := r.rdb.Get(r.ctx, "track:"+id).Result()
	if err == nil {
		var track models.Track
		json.Unmarshal([]byte(cachedTrack), &track)
		return &track, nil
	}

	query := `SELECT id, title, artist, duration FROM tracks WHERE id = $1`
	var track models.Track
	err = r.db.QueryRow(query, id).Scan(&track.ID, &track.Title, &track.Artist, &track.Duration)
	if err != nil {
		return nil, err
	}

	trackData, _ := json.Marshal(track)
	r.rdb.Set(r.ctx, "track:"+id, trackData, 5*time.Minute)

	return &track, nil
}