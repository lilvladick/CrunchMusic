package service

import (
	"CrunchServer/postgres"
	"database/sql"
)

func FetchAllTracks(db *sql.DB) ([]postgres.Track, error) {
	var tracks []postgres.Track
	rows, err := postgres.GetTracks(db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var track postgres.Track
		err = rows.Scan(&track.ID, &track.Title, &track.Genre, &track.Duration)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}
