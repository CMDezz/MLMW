package queries

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func (query Query) UpsertTracksPlayListsQuery(ctx *gin.Context, req []models.TrackPlayListSchema, trackId int64) ([]models.TrackPlayListSchema, error) {
	// SQL insert dynamic values

	var values []interface{}
	var placeholders []string

	for i, tp := range req {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))

		values = append(values, tp.TrackId, tp.PlaylistId)
	}

	//SQL delete rows not in the req
	var sqlDelete string
	if len(placeholders) == 0 {

		sqlDelete = fmt.Sprintf(`
			DELETE FROM %s 
			WHERE track_id = %d 
		`, utils.TABLE_TRACKPLAYLIST, trackId)
	} else {
		sqlDelete = fmt.Sprintf(`
			DELETE FROM %s 
			WHERE track_id = %d 
			AND (track_id, playlist_id) NOT IN (%s)
		`, utils.TABLE_TRACKPLAYLIST, trackId, strings.Join(placeholders, ", "))
	}

	// Execute the DELETE query
	_, err := query.store.Exec(sqlDelete, values...)
	if err != nil {
		return []models.TrackPlayListSchema{}, err
	}

	// Dynamic sql insert
	if len(req) == 0 { // no insert if there is no items
		return []models.TrackPlayListSchema{}, nil
	}
	sqlCmd := fmt.Sprintf("INSERT INTO %s (track_id, playlist_id) VALUES", utils.TABLE_TRACKPLAYLIST)
	sqlCmd += strings.Join(placeholders, ", ")
	sqlCmd += " ON CONFLICT (track_id, playlist_id) DO NOTHING RETURNING *"

	var res []models.TrackPlayListSchema

	err = query.store.Select(&res, sqlCmd, values...)
	if err != nil {
		return []models.TrackPlayListSchema{}, err
	}

	return res, nil
}

func (query Query) GetDataTrackPlaylistByIdQuery(ctx *gin.Context, id int64) ([]models.TrackPlayListSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE track_id=$1", utils.TABLE_TRACKPLAYLIST)

	var res []models.TrackPlayListSchema
	err := query.store.Select(&res, sqlCmd, id)
	if err != nil {
		return []models.TrackPlayListSchema{}, nil
	}
	return res, nil
}

func (query Query) GetDataPlaylistHasTracksQuery(ctx *gin.Context, id int64) ([]models.TrackSchema, error) {
	//get from join table
	sqlCmd := fmt.Sprintf(`
	SELECT t.*
	FROM %s tp
	JOIN %s t ON tp.track_id = t.id
	WHERE tp.playlist_id = $1
`, utils.TABLE_TRACKPLAYLIST, utils.TABLE_TRACKS)

	var tracks []models.TrackSchema
	err := query.store.Select(&tracks, sqlCmd, id)
	if err != nil {
		return []models.TrackSchema{}, err
	}

	//relfect to track table
	return tracks, nil
}
