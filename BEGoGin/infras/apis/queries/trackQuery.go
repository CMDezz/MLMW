package queries

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (query Query) GetAllPublicTrackQuery(ctx *gin.Context) ([]models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE is_public=$1 AND is_deleted=$2", utils.TABLE_TRACKS)

	var res []models.TrackSchema

	err := query.store.Select(&res, sqlCmd, true, false)

	if err != nil {
		return []models.TrackSchema{}, err
	}

	return res, nil
}

func (query Query) GetAllTracksByUserIdQuery(ctx *gin.Context, id int64) ([]models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND is_deleted=$2", utils.TABLE_TRACKS)

	var res []models.TrackSchema

	err := query.store.Select(&res, sqlCmd, id, false)

	if err != nil {
		return []models.TrackSchema{}, err
	}

	return res, nil
}

func (query Query) GetTrackById(ctx *gin.Context, id int64) (models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND is_deleted=$2 LIMIT 1", utils.TABLE_TRACKS)

	var res models.TrackSchema

	err := query.store.Get(&res, sqlCmd, id, false)

	if err != nil {
		return models.TrackSchema{}, err
	}

	return res, nil
}

func (query Query) CreateTrackQuery(ctx *gin.Context, track models.TrackSchema) (models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf("INSERT INTO %s (title, artist, album,genre,release_year,user_id,url,cover_image) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *", utils.TABLE_TRACKS)

	var res models.TrackSchema

	// playlistId := utils.MarshalGetString(track.PlaylistId)

	err := query.store.QueryRowxContext(ctx, sqlCmd, track.Title, track.Artist, track.Album, track.Genre, track.ReleaseYear, track.UserId, track.Url, track.CoverImage).StructScan(&res)

	if err != nil {
		return models.TrackSchema{}, err
	}
	return res, nil
}

func (query Query) UpdateTrackQuery(ctx *gin.Context, track models.TrackSchema) (models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf(`
		UPDATE %s
        SET title = $2, artist = $3, album = $4, genre = $5, release_year = $6, url = $7, is_public = $8, cover_image = $9
        WHERE id = $1
        RETURNING *
	`, utils.TABLE_TRACKS)

	var res models.TrackSchema
	err := query.store.QueryRowxContext(ctx, sqlCmd, track.Id, track.Title, track.Artist, track.Album, track.Genre, track.ReleaseYear, track.Url, track.IsPublic, track.CoverImage).StructScan(&res)

	if err != nil {
		return models.TrackSchema{}, err
	}
	return res, nil
}

func (query Query) DeleteTrackByIdQuery(ctx *gin.Context, id int64) error {
	sqlCmd := fmt.Sprintf(`
		UPDATE %s
        SET is_deleted = $2, is_public = $3
        WHERE id = $1
        RETURNING *
	`, utils.TABLE_TRACKS)

	_, err := query.store.Exec(sqlCmd, id, true, false)

	if err != nil {
		return err
	}
	return nil
}

func (query Query) SearchingTracksQuery(ctx *gin.Context, keyword string) ([]models.TrackSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE (title ILIKE $1 OR artist ILIKE $1) AND is_deleted=$2", utils.TABLE_TRACKS)

	var res []models.TrackSchema

	err := query.store.Select(&res, sqlCmd, keyword, false)

	if err != nil {
		return []models.TrackSchema{}, err
	}

	return res, nil
}
