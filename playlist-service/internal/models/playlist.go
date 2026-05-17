package models

type Playlist struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	UserID   string   `json:"user_id"`
	TrackIDs []string `json:"track_ids"`
}