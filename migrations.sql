CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    password BYTEA NOT NULL
);

CREATE TABLE tracks (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    filepath TEXT NOT NULL,
    user_id int NOT NULL,
    genre VARCHAR NOT NULL,
    duration TIME
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE likes (
    user_id int NOT NULL,
    track_id int NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_track_id FOREIGN KEY (track_id) REFERENCES tracks (id),
    CONSTRAINT unique_like_pair UNIQUE (user_id, track_id)
);