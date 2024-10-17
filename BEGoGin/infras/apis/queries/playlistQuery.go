package queries

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (query Query) GetAllPublicPlaylistsQuery(ctx *gin.Context) ([]models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE is_public=$1 AND is_deleted=$2", utils.TABLE_PLAYLISTS)

	var res []models.PlaylistSchema

	err := query.store.Select(&res, sqlCmd, true, false)

	if err != nil {
		return []models.PlaylistSchema{}, err
	}

	return res, nil
}

func (query Query) GetAllPlaylistsByUserIdQuery(ctx *gin.Context, id int64) ([]models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND is_deleted=$2", utils.TABLE_PLAYLISTS)

	var res []models.PlaylistSchema

	err := query.store.Select(&res, sqlCmd, id, false)

	if err != nil {
		return []models.PlaylistSchema{}, err
	}

	return res, nil
}

func (query Query) CreatePlaylistQuery(ctx *gin.Context, playlist models.PlaylistSchema) (models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf("INSERT INTO %s (playlist_name, description, cover_image,user_id) VALUES($1,$2,$3,$4) RETURNING *", utils.TABLE_PLAYLISTS)

	var res models.PlaylistSchema

	// playlistId := utils.MarshalGetString(track.PlaylistId)

	err := query.store.QueryRowxContext(ctx, sqlCmd, playlist.PlaylistName, playlist.Description, playlist.CoverImage, playlist.UserId).StructScan(&res)

	if err != nil {
		return models.PlaylistSchema{}, err
	}
	return res, nil
}

func (query Query) GetPlaylistById(ctx *gin.Context, id int64) (models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND is_deleted=$2 LIMIT 1", utils.TABLE_PLAYLISTS)

	var res models.PlaylistSchema

	err := query.store.Get(&res, sqlCmd, id, false)

	if err != nil {
		return models.PlaylistSchema{}, err
	}

	return res, nil
}

func (query Query) UpdatePlaylistQuery(ctx *gin.Context, playlist models.PlaylistSchema) (models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf(`
		UPDATE %s
        SET playlist_name = $2, description = $3, cover_image = $4, is_public = $5
        WHERE id = $1
        RETURNING *
	`, utils.TABLE_PLAYLISTS)

	var res models.PlaylistSchema
	err := query.store.QueryRowxContext(ctx, sqlCmd, playlist.Id, playlist.PlaylistName, playlist.Description, playlist.CoverImage, playlist.IsPublic).StructScan(&res)

	if err != nil {
		return models.PlaylistSchema{}, err
	}
	return res, nil
}

func (query Query) DeletePlaylistByIdQuery(ctx *gin.Context, id int64) error {
	sqlCmd := fmt.Sprintf(`
		UPDATE %s
        SET is_deleted = $2, is_public = $3
        WHERE id = $1
        RETURNING *
	`, utils.TABLE_PLAYLISTS)

	_, err := query.store.Exec(sqlCmd, id, true, false)

	if err != nil {
		return err
	}
	return nil
}

func (query Query) SearchingPlaylistsQuery(ctx *gin.Context, keyword string) ([]models.PlaylistSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE playlist_name ILIKE $1 AND is_deleted=$2", utils.TABLE_PLAYLISTS)

	var res []models.PlaylistSchema

	err := query.store.Select(&res, sqlCmd, keyword, false)

	if err != nil {
		return []models.PlaylistSchema{}, err
	}

	return res, nil
}
