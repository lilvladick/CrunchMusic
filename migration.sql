CREATE TABLE "users" (
    "id" BIGSERIAL NOT NULL UNIQUE,
    "name" VARCHAR NOT NULL,
    "login" VARCHAR NOT NULL UNIQUE,
    "password" BYTEA NOT NULL,
    PRIMARY KEY("id")
);

CREATE TABLE "tracks" (
    "id" BIGSERIAL NOT NULL UNIQUE,
    "title" VARCHAR NOT NULL,
    "user_id" BIGINT NOT NULL,
    "genre" VARCHAR NOT NULL,
    "duration" INTERVAL,
    "file_path" VARCHAR NOT NULL,  -- Путь к файлу на сервере
    PRIMARY KEY("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id")
    ON UPDATE NO ACTION ON DELETE CASCADE
);
