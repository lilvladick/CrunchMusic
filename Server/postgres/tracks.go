package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func UploadTrack(title string, filepath string, userID int, genre string, duration pq.NullTime) error {
	sqlStatement := `
        INSERT INTO tracks (title, filepath, user_id, genre, duration)
        VALUES ($1, $2, $3, $4, $5);
    `
	_, err := db.Exec(sqlStatement, title, filepath, userID, genre, duration)
	return err
}

func DeleteTrack(id int) error {
	sqlStatement := `
        DELETE FROM tracks
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func GetTrackByID(id int) (*Track, error) {
	sqlStatement := `
        SELECT * FROM tracks
        WHERE id = $1;
    `

	row := db.QueryRow(sqlStatement, id)
	var track Track

	err := row.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with ID %v", id)
		}
		return nil, fmt.Errorf("error scanning user row: %v", err)
	}

	return &track, nil
}

func GetTrackByTitle(title string) ([]Track, error) {
	sqlStatement := `
        SELECT * FROM tracks
        WHERE title = $1;
    `
	rows, err := db.Query(sqlStatement, title)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var tracks []Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("error scanning track row: %v", err)
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tracks rows: %v", err)
	}

	return tracks, nil
}

func GetTrackByGenre(genre string) ([]Track, error) {
	sqlStatement := "SELECT * FROM tracks WHERE genre = $1"

	rows, err := db.Query(sqlStatement, genre)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var tracks []Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("error scanning track row: %v", err)
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tracks rows: %v", err)
	}

	return tracks, nil
}

func GetTracksFromPlaylist(playlist_id int) ([]Track, error) {
	sqlStatement := "SELECT t.* FROM tracks t JOIN playlist_tracks pt ON t.id = pt.track_id WHERE pt.playlist_id = $1; "

	rows, err := db.Query(sqlStatement, playlist_id)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var tracks []Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("error scanning track row: %v", err)
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tracks rows: %v", err)
	}

	return tracks, nil
}

func GetTracks() ([]Track, error) {
	sqlStatement := `
        SELECT * FROM tracks;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var tracks []Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("error scanning track row: %v", err)
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tracks rows: %v", err)
	}

	return tracks, nil
}

func Get100Tracks() ([]Track, error) {
	sqlStatement := "SELECT * FROM tracks LIMIT 100;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var tracks []Track
	for rows.Next() {
		var track Track
		err := rows.Scan(&track.ID, &track.Title, &track.Filepath, &track.UserID, &track.Genre, &track.Duration)
		if err != nil {
			return nil, fmt.Errorf("error scanning track row: %v", err)
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over tracks rows: %v", err)
	}

	return tracks, nil
}
