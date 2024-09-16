CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    password BYTEA NOT NULL
);

CREATE TABLE tracks (
    id UUID PRIMARY KEY,
    title VARCHAR NOT NULL,
    filepath TEXT NOT NULL,
    user UUID NOT NULL,
    genre VARCHAR NOT NULL,
    duration TIMESTAMP
);

CREATE TABLE likes (
    user_id UUID NOT NULL,
    track_id UUID NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_track_id FOREIGN KEY (track_id) REFERENCES tracks (id)
);

ALTER TABLE tracks ADD CONSTRAINT fk_user FOREIGN KEY (user) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE NO ACTION;