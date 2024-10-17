package handlers

import (
	"MLMW/BEGoGin/models"

	"github.com/gin-gonic/gin"
)

func (handler Handler) SearchingHandler(ctx *gin.Context, keyword string) (models.SearchingResponse, error) {
	//for partial matches
	searchKeyword := "%" + keyword + "%"
	resTracks, err := handler.query.SearchingTracksQuery(ctx, searchKeyword)
	if err != nil {
		return models.SearchingResponse{}, err
	}

	resPlaylists, err := handler.query.SearchingPlaylistsQuery(ctx, searchKeyword)
	if err != nil {
		return models.SearchingResponse{}, err
	}
	return models.SearchingResponse{
		Tracks:    mapTrackSchemaToTrackResponse(&resTracks),
		Playlists: mapPlaylistSchemaToPlaylistResponse(&resPlaylists),
	}, nil
}
