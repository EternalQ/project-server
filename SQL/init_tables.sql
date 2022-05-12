CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" text UNIQUE NOT NULL,
    "encrypted_password" text NOT NULL,
    "created_at" timestamp NOT NULL
);

CREATE TABLE "posts" (
    "id" BIGSERIAL PRIMARY KEY,
    "text" text,
    "created_at" timestamp NOT NULL,
    "photo_url" text,
    "user_id" bigint NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "post_tags" (
    "post_id" bigint NOT NULL,
    "tag" text NOT NULL,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE
);

CREATE TABLE "album" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" text,
    "created_at" timestamp NOT NULL,
    "user_id" bigint NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE TABLE "album_photos" (
    "photo_url" text NOT NULL,
    "album_id" bigint NOT NULL,
    -- FOREIGN KEY ("photo_url") REFERENCES "photos" ("url") on DELETE CASCADE,
    FOREIGN KEY ("album_id") REFERENCES "album" ("id") ON DELETE CASCADE
);

CREATE TABLE "comments" (
    "id" BIGSERIAL PRIMARY KEY,
    "comment" text,
    "created_at" timestamp NOT NULL,
    "post_id" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);