CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" text UNIQUE NOT NULL,
    "encrypted_password" text NOT NULL,
    "created_at" timestamp
);

CREATE TABLE "photos" (
    "id" BIGSERIAL PRIMARY KEY,
    "url" text NOT NULL
);

CREATE TABLE "posts" (
    "id" BIGSERIAL PRIMARY KEY,
    "text" text,
    "created_at" timestamp NOT NULL,
    "photo_id" bigint,
    "user_id" bigint NOT NULL,
    FOREIGN KEY("photo_id") REFERENCES "photos" ("id"),
    FOREIGN KEY("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "tags" (
    "id" BIGSERIAL PRIMARY KEY,
    "tag" text NOT NULL UNIQUE
);

CREATE TABLE "post_tags" (
    "post_id" bigint NOT NULL,
    "tag_id" bigint NOT NULL,
    FOREIGN KEY("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
    FOREIGN KEY("tag_id") REFERENCES "tags" ("id")
);

CREATE TABLE "albums" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "created_at" timestamp NOT NULL,
    "user_id" bigint NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "album_photos" (
    "photo_id" bigint NOT NULL,
    "album_id" bigint NOT NULL,
    FOREIGN KEY("album_id") REFERENCES "albums" ("id") ON DELETE CASCADE,
    FOREIGN KEY("photo_id") REFERENCES "photos" ("id")
);

CREATE TABLE "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "comment" text,
    "created_at" timestamp NOT NULL,
    "post_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    FOREIGN KEY("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    FOREIGN KEY("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE
);