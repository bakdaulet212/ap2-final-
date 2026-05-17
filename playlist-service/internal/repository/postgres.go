package repository

import (
	"database/sql"
	"log"

	"github.com/bakdaulet212/ap2-final-/playlist-service/internal/models"
	_ "github.com/lib/pq"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *PlaylistRepository {
	connStr := "host=localhost port=5432 user=user_admin password=my_secret_password dbname=music_app_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS playlists (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title VARCHAR(100) NOT NULL,
		user_id UUID NOT NULL
	);
	CREATE TABLE IF NOT EXISTS playlist_tracks (
		playlist_id UUID REFERENCES playlists(id) ON DELETE CASCADE,
		track_id UUID NOT NULL,
		PRIMARY KEY (playlist_id, track_id)
	);`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	return &PlaylistRepository{db: db}
}

func (r *PlaylistRepository) CreatePlaylist(playlist *models.Playlist) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var playlistID string
	query := `INSERT INTO playlists (title, user_id) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(query, playlist.Title, playlist.UserID).Scan(&playlistID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	trackQuery := `INSERT INTO playlist_tracks (playlist_id, track_id) VALUES ($1, $2)`
	for _, trackID := range playlist.TrackIDs {
		_, err = tx.Exec(trackQuery, playlistID, trackID)
		if err != nil {
			tx.Rollback()
			return "", err
		}
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return playlistID, nil
}