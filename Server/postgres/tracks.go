package postgres

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func UploadTrack(db *sql.DB, id int, title string, filepath string, user_id string, genre string, now time.Time) (string, error) {
	sqlStatement := `
        INSERT INTO tracks (id, title, filepath, user_id, genre, now)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `
	var trackId string
	err := db.QueryRow(sqlStatement, id, title, filepath, user_id, genre, now).Scan(&trackId)
	return trackId, err
}

func DeleteTrack(db *sql.DB, id int) error {
	sqlStatement := `
        DELETE FROM tracks
        WHERE id = $1;
    `
	_, err := db.Exec(sqlStatement, id)
	return err
}

func GetTrackByID(db *sql.DB, id int) (*sql.Row, error) {
	sqlStatement := `
        SELECT * FROM tracks
        WHERE id = $1;
    `
	return db.QueryRow(sqlStatement, id), nil
}

func GetTrackByTitle(db *sql.DB, title string) (*sql.Row, error) {
	sqlStatement := `
        SELECT * FROM tracks
        WHERE title = $1;
    `
	return db.QueryRow(sqlStatement, title), nil
}

func GetTrackByGenre(db *sql.DB, genre string) (*sql.Rows, error) {
	sqlStatement := `
        SELECT * FROM tracks
        WHERE Genre = $1;
    `
	return db.Query(sqlStatement, genre)
}

func GetTracks(db *sql.DB) (*sql.Rows, error) {
	sqlStatement := `
        SELECT * FROM tracks;
    `
	return db.Query(sqlStatement)
}
