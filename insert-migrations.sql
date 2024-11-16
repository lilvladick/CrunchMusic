
INSERT INTO users (name, login, password)
VALUES
    ('John Doe', 'johndoe', 'password123'),
    ('Jane Smith', 'janesmith', 'password456'),
    ('Bob Johnson', 'bobjohnson', 'password789');

INSERT INTO tracks (title, filepath, user_id, genre, duration)
VALUES
    ('Track 1', '/Server/music/track1.mp3', 1, 'Rock', '00:03:30'),
    ('Track 2', '/Server/music/track2.mp3', 2, 'Pop', '00:04:15'),
    ('Track 3', '/Server/music/track3.mp3', 3, 'Jazz', '00:05:00');

INSERT INTO likes (user_id, track_id)
VALUES
    (1, 1),  
    (1, 2),  
    (2, 1),  
    (2, 3),  
    (3, 2);  