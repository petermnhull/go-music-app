package models

import "time"

// User holds information for music listener
type User struct {
	ID              int64     `json:"id"`
	SpotifyUsername string    `json:"spotify_username"`
	LastfmUsername  string    `json:"lastfm_username"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
