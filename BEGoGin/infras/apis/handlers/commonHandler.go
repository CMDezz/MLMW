package handlers

import (
	"MLMW/BEGoGin/models"

	"github.com/gin-gonic/gin"
)

func mapTrackPlaylistResponse(list *[]models.TrackPlayListSchema) []models.TracksPlaylistsResponse {
	res := make([]models.TracksPlaylistsResponse, len(*list))
	for i, item := range *list {
		res[i] = models.TracksPlaylistsResponse{
			TrackId:    item.TrackId,
			PlaylistId: item.PlaylistId,
		}
	}
	return res
}

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

func (handler Handler) UpsertTracksPlayListsHandler(ctx *gin.Context, req models.UpsertTracksPlaylistsModel) ([]models.TracksPlaylistsResponse, error) {

	newTrackPlaylists := make([]models.TrackPlayListSchema, len(req.PlaylistId))
	for i, playlistId := range req.PlaylistId {
		newTrackPlaylists[i] = models.TrackPlayListSchema{
			TrackId:    req.TrackId,
			PlaylistId: playlistId,
		}
	}

	res, err := handler.query.UpsertTracksPlayListsQuery(ctx, newTrackPlaylists, req.TrackId)
	if err != nil {
		return []models.TracksPlaylistsResponse{}, err
	}

	return mapTrackPlaylistResponse(&res), nil

}

func (handler Handler) GetDataTrackPlaylistByIdHandler(ctx *gin.Context, id int64) (models.DetailTrackPlaylistResponse, error) {

	res, err := handler.query.GetDataTrackPlaylistByIdQuery(ctx, id)

	if err != nil {
		return models.DetailTrackPlaylistResponse{}, err
	}

	var resId []int64
	for _, trackPl := range res {
		resId = append(resId, trackPl.PlaylistId)
	}

	return models.DetailTrackPlaylistResponse{
		PlaylistId: resId,
	}, nil

}

func (handler Handler) GetFullPlaylistDetailHandler(ctx *gin.Context, id int64) (models.PlaylistFullDetailResponse, error) {
	// get playlist detail
	playlist, err := handler.query.GetPlaylistById(ctx, id)
	if err != nil {
		return models.PlaylistFullDetailResponse{}, nil
	}
	// get tracks belong to playlit
	tracks, err := handler.query.GetDataPlaylistHasTracksQuery(ctx, id)
	if err != nil {
		return models.PlaylistFullDetailResponse{}, nil
	}
	// combine response

	res := models.PlaylistFullDetailResponse{
		Playlist: models.PlaylistResponse{
			Id:           playlist.Id,
			PlaylistName: playlist.PlaylistName,
			Description:  playlist.Description,
			CoverImage:   playlist.CoverImage,
			IsPublic:     playlist.IsPublic,
			UserId:       playlist.UserId,
			CreatedAt:    playlist.CreatedAt,
		},
		Tracks: mapTrackSchemaToTrackResponse(&tracks),
	}

	return res, nil

}
