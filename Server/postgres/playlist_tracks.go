package postgres

import (
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func GetPlaylistTracksByPlaylistID(playlist_id int) ([]Playlist_tracks, error) {
	sqlStatement := `
        SELECT * FROM playlist_tracks
        WHERE playlist_id = $1;
    `

	rows, err := db.Query(sqlStatement, playlist_id)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists_tracks []Playlist_tracks
	for rows.Next() {
		var playlist_track Playlist_tracks
		err := rows.Scan(&playlist_track.PlaylistID, &playlist_track.TrackID, &playlist_track.AddedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist_tracks row: %v", err)
		}
		playlists_tracks = append(playlists_tracks, playlist_track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists_tracks, nil
}

func GetPlaylistTracksByTrackID(track_id int) ([]Playlist_tracks, error) {
	sqlStatement := `
        SELECT * FROM playlist_tracks
        WHERE track_id = $1;
    `

	rows, err := db.Query(sqlStatement, track_id)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists_tracks []Playlist_tracks
	for rows.Next() {
		var playlist_track Playlist_tracks
		err := rows.Scan(&playlist_track.PlaylistID, &playlist_track.TrackID, &playlist_track.AddedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist_tracks row: %v", err)
		}
		playlists_tracks = append(playlists_tracks, playlist_track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists_tracks, nil
}

func GetPlaylistTracks() ([]Playlist_tracks, error) {
	sqlStatement := `
        SELECT * FROM playlist_tracks`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, fmt.Errorf("error querying tracks: %v", err)
	}
	defer rows.Close()

	var playlists_tracks []Playlist_tracks
	for rows.Next() {
		var playlist_track Playlist_tracks
		err := rows.Scan(&playlist_track.PlaylistID, &playlist_track.TrackID, &playlist_track.AddedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist_tracks row: %v", err)
		}
		playlists_tracks = append(playlists_tracks, playlist_track)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over playlists rows: %v", err)
	}

	return playlists_tracks, nil
}

func AddTrackToPlaylist(playlist_id int, track_id int, added_at pq.NullTime) error {
	sqlStatement := `
        INSERT INTO playlist_tracks (playlist_id,track_id,added_at)
        VALUES ($1, $2,$3);
    `
	_, err := db.Exec(sqlStatement, playlist_id, track_id, added_at)
	return err
}
