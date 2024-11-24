package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPlaylistByID(id int) (*Playlist, error) {
	sqlStatement := `
        SELECT * FROM playlists
        WHERE id = $1;
    `

	row := db.QueryRow(sqlStatement, id)
	var playlist Playlist

	err := row.Scan(&playlist.ID, &playlist.Name, &playlist.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("playlist not found with ID %v", id)
		}
		return nil, fmt.Errorf("error scanning playlist row: %v", err)
	}

	return &playlist, nil
}

func GetPlaylists() ([]Playlist, error) {
	sqlStatement := `
        SELECT * FROM playlists;
    `
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists []Playlist
	for rows.Next() {
		var playlist Playlist
		err := rows.Scan(&playlist.ID, &playlist.Name, &playlist.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist row: %v", err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists, nil
}

func GetPlaylistByName(name string) ([]Playlist, error) {
	sqlStatement := `
        SELECT * FROM playlists
        WHERE name = $1;
    `

	rows, err := db.Query(sqlStatement, name)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists []Playlist
	for rows.Next() {
		var playlist Playlist
		err := rows.Scan(&playlist.ID, &playlist.Name, &playlist.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist row: %v", err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists, nil
}

func GetPlaylistByUserID(user_id int) ([]Playlist, error) {
	sqlStatement := `
        SELECT * FROM playlists
        WHERE id = $1;
    `

	rows, err := db.Query(sqlStatement, user_id)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists []Playlist
	for rows.Next() {
		var playlist Playlist
		err := rows.Scan(&playlist.ID, &playlist.Name, &playlist.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist row: %v", err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists, nil
}

func CreatePlaylist(name string, userID int) error {
	sqlStatement := `
        INSERT INTO playlists (name,user_id)
        VALUES ($1, $2);
    `
	_, err := db.Exec(sqlStatement, name, userID)
	return err
}
