package models

import "time"

//USER APIS
type CreateUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	// Token     string    `json:"token"`
	// ExpiredAt time.Time `json:"expired_at"`
}

type LoginUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

//skip isdeleted
type TrackResponse struct {
	Id          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	CoverImage  string    `json:"cover_image" db:"cover_image"`
	Artist      string    `json:"artist" db:"artist"`
	Album       string    `json:"album" db:"album"`
	Genre       string    `json:"genre" db:"genre"`
	ReleaseYear int       `json:"release_year" db:"release_year"`
	UserId      int64     `json:"user_id" db:"user_id"`
	Url         string    `json:"url" db:"url"`
	IsPublic    bool      `form:"is_public" db:"is_public"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type DeletedResponse struct {
	Message string `json:"id"`
}

// THIS 2 WILL BE REPLACED IN THE FUTURE IF HAVE TIME
type GetAllPublicTracksResponse struct {
	Tracks []TrackResponse
}
type GetAllTracksResponse struct {
	Tracks []TrackResponse
}

//TODO : REPLACE THIS TYPE FOR 2 TYPE ABOVE
type ListTrackResponse struct {
	Tracks []TrackResponse
}

type UpdateTrackResponse struct {
	TrackResponse
}

type PlaylistResponse struct {
	Id           int64     `json:"id" db:"id"`
	PlaylistName string    `json:"playlist_name" db:"playlist_name"`
	Description  string    `json:"description" db:"description"`
	CoverImage   string    `json:"cover_image" db:"cover_image"`
	IsPublic     bool      `json:"is_public" db:"is_public"`
	UserId       int64     `json:"user_id" db:"user_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type ListPlaylistResponse struct {
	Playlists []PlaylistResponse
}

//common
type SearchingResponse struct {
	Tracks    []TrackResponse
	Playlists []PlaylistResponse
}
