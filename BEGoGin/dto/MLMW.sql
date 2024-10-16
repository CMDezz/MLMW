CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hased_password" varchar NOT NULL,
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
  "title" varchar(100) NOT NULL,
  "artist" varchar(100) NOT NULL,
  "album" varchar(100) NOT NULL,
  "genre" varchar(100) NOT NULL,
  "release_year" int NOT NULL,
  "playlist_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "is_public" bool NOT NULL DEFAULT true,
  "is_deleted" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "playlists" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tracks" ADD FOREIGN KEY ("playlist_id") REFERENCES "playlists" ("id");

ALTER TABLE "tracks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
