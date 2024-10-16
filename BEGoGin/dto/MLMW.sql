CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "full_name" varchar,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "playlists" (
  "id" bigserial PRIMARY KEY,
  "playlist_name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "cover_image" varchar NOT NULL,
  "user_id" bigserial NOT NULL,
  "is_public" bool NOT NULL DEFAULT true,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tracks" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "artist" varchar NOT NULL,
  "album" varchar NOT NULL,
  "genre" varchar NOT NULL,
  "release_year" int NOT NULL,
  "user_id" bigserial NOT NULL,
  "url" varchar NOT NULL,
  "is_public" bool NOT NULL DEFAULT true,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "track_playlist" (
  "track_id" BIGINT,
  "playlist_id" BIGINT,
  PRIMARY KEY ("track_id", "playlist_id")
);

ALTER TABLE "playlists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tracks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "track_playlist" ADD FOREIGN KEY ("track_id") REFERENCES "tracks" ("id");

ALTER TABLE "track_playlist" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");
