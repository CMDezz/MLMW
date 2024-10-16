package models

import "time"

type UserSchema struct {
	Id             int64     `json:"id" db:"id"`
	Username       string    `json:"username" db:"username"`
	FullName       string    `json:"full_name" db:"full_name"`
	Email          string    `json:"email" db:"email"`
	HashedPassword string    `json:"hashed_password" db:"hashed_password"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type TrackSchema struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Artist      string    `json:"artist" db:"artist"`
	Album       string    `json:"album" db:"album"`
	Genre       string    `json:"genre" db:"genre"`
	ReleaseYear int       `json:"release_year" db:"release_year"`
	Duration    int       `json:"duration" db:"duration"`
	UserId      int64     `json:"user_id" db:"user_id"`
	Url         string    `json:"url" db:"url"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
	IsDeleted   bool      `json:"is_deleted" db:"is_deleted"`
	CoverImage  string    `json:"cover_image" db:"cover_image"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type TrackPlayListSchema struct {
	TrackId    int64 `json:"track_id" db:"track_id"`
	PlaylistId int64 `json:"playlist_id" db:"playlist_id"`
}

type PlaylistSchema struct {
	Id           int64     `json:"id" db:"id"`
	PlaylistName string    `json:"playlist_name" db:"playlist_name"`
	Description  string    `json:"description" db:"description"`
	CoverImage   string    `json:"cover_image" db:"cover_image"`
	UserId       int64     `json:"user_id" db:"user_id"`
	IsPublic     bool      `json:"is_public" db:"is_public"`
	IsDeleted    bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
