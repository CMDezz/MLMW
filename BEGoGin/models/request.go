package models

import "mime/multipart"

// REQUEST APIS MODELS

//COMMON
type ByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
type SearchingRequest struct {
	Keyword string `form:"keyword" binding:"required,min=1"`
}

//USER APIS
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

//TRACK APIS

type CreateTrackFormRequest struct { // Upload by JS FormData, parse to this struct later
	Title       string `form:"title" binding:"required"`
	Artist      string `form:"artist" binding:"required"`
	Album       string `form:"album" binding:"required"`
	Genre       string `form:"genre" binding:"required"`
	ReleaseYear int    `form:"release_year" binding:"required"`
	Duration    int    `form:"duration" binding:"required"`
	// UserId      int64                 `form:"user_id" binding:"required"`
	CoverImage *multipart.FileHeader `form:"cover_image" binding:"required"`
	TrackFile  *multipart.FileHeader `form:"track_file" binding:"required"`
}

type UpdateTrackFormRequest struct { // Upload by JS FormData, parse to this struct later
	Id          int64                 `form:"id"  binding:"required"`
	Title       string                `form:"title" binding:"required"`
	Artist      string                `form:"artist" binding:"required"`
	Album       string                `form:"album" binding:"required"`
	Genre       string                `form:"genre" binding:"required"`
	ReleaseYear int                   `form:"release_year" binding:"required"`
	Duration    int                   `form:"duration" binding:"required"`
	IsPublic    bool                  `form:"is_public"`
	CoverImage  *multipart.FileHeader `form:"cover_image"`
	TrackFile   *multipart.FileHeader `form:"track_file"`
}

//PLAYLIST APIS

type CreatePlaylistFormRequest struct {
	PlaylistName string `form:"playlist_name" binding:"required"`
	// UserId       int64                 `form:"user_id" binding:"required"`
	Description string                `form:"description" binding:"required"`
	CoverImage  *multipart.FileHeader `form:"cover_image" `
}

type UpdatePlaylistFormRequest struct {
	Id           int64                 `form:"id"  binding:"required"`
	PlaylistName string                `form:"playlist_name" binding:"required"`
	Description  string                `form:"description" binding:"required"`
	CoverImage   *multipart.FileHeader `form:"cover_image" `
	IsPublic     bool                  `form:"is_public"`
}
