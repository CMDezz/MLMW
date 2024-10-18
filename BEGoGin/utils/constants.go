package utils

import "time"

// constant config path
const CONFIG_PATH = "."

const MIN_LENGTH_SECRET_KEY = 12
const DEFAULT_ACCESS_TOKEN_DURATION = time.Duration(1800) * time.Second

const AUTHORIZATION_HEADER_KEY = "authorization"
const AUTHORIZATION_BEAER_TYPE = "bearer"
const AUTHORIZATION_PAYLOAD_KEY = "authorization_payload"

const TABLE_USERS = "users"
const TABLE_TRACKS = "tracks"
const TABLE_PLAYLISTS = "playlists"
const TABLE_TRACKPLAYLIST = "track_playlist"
const UPLOAD_DIR_TRACK = "upload/tracks"
const UPLOAD_DIR_IMAGE = "upload/images"
