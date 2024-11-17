CREATE TABLE playlists (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    user_id int NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE playlist_tracks (
    playlist_id int NOT NULL,
    track_id int NOT NULL,
    added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_playlist_id FOREIGN KEY (playlist_id) REFERENCES playlists (id),
    CONSTRAINT fk_track_id FOREIGN KEY (track_id) REFERENCES tracks (id),
    CONSTRAINT unique_playlist_track_pair UNIQUE (playlist_id, track_id)
);

INSERT INTO playlists ( name, user_id)
VALUES
    ( 'My Favorite Tracks', 1),
    ( 'Workout Playlist', 1),
    ( 'Relaxation Music', 2),
    ( 'Party Mix', 3);

INSERT INTO playlist_tracks (playlist_id, track_id, added_at)
VALUES
    (1, 1, now()),  
    (1, 2, now()),  
    (1, 3, now()),  
    (2, 1, now()),  
    (3, 3, now()),  
    (4, 2, now()),  