package handlers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

func mapPlaylistSchemaToPlaylistResponse(playlistsSchema *[]models.PlaylistSchema) []models.PlaylistResponse {
	playlistsResponse := make([]models.PlaylistResponse, len(*playlistsSchema))
	for i, playlist := range *playlistsSchema {
		playlistsResponse[i] = models.PlaylistResponse{
			Id:           playlist.Id,
			IsPublic:     playlist.IsPublic,
			UserId:       playlist.UserId,
			PlaylistName: playlist.PlaylistName,
			CreatedAt:    playlist.CreatedAt,
			Description:  playlist.Description,
			CoverImage:   playlist.CoverImage,
		}
	}
	return playlistsResponse
}

func (handler Handler) GetAllPublicPlaylistsHandler(ctx *gin.Context) (models.ListPlaylistResponse, error) {
	res, err := handler.query.GetAllPublicPlaylistsQuery(ctx)

	if err != nil {
		return models.ListPlaylistResponse{}, err
	}

	return models.ListPlaylistResponse{
		Playlists: mapPlaylistSchemaToPlaylistResponse(&res),
	}, nil
}

func (handler Handler) GetAllPlaylistsByUserIdHandler(ctx *gin.Context, id int64) (models.ListPlaylistResponse, error) {
	res, err := handler.query.GetAllPlaylistsByUserIdQuery(ctx, id)

	if err != nil {
		return models.ListPlaylistResponse{}, err
	}

	return models.ListPlaylistResponse{
		Playlists: mapPlaylistSchemaToPlaylistResponse(&res),
	}, nil
}

func (handler Handler) CreatePlaylistHandler(ctx *gin.Context, req models.CreatePlaylistFormRequest, userId int64) (models.PlaylistResponse, error) {

	//Turn binaryfile into osFile and save to upload Folder
	filePath, err := utils.SaveUploadedFile(req.CoverImage, utils.UPLOAD_DIR_IMAGE)
	if err != nil {
		return models.PlaylistResponse{}, err
	}

	playlist := models.PlaylistSchema{
		PlaylistName: req.PlaylistName,
		Description:  req.Description,
		UserId:       userId,
		CoverImage:   filePath,
	}

	res, err := handler.query.CreatePlaylistQuery(ctx, playlist)
	if err != nil {
		return models.PlaylistResponse{}, err
	}
	return models.PlaylistResponse{
		Id:           res.Id,
		PlaylistName: res.PlaylistName,
		Description:  res.Description,
		CoverImage:   res.CoverImage,
		IsPublic:     res.IsPublic,
		UserId:       res.UserId,
		CreatedAt:    res.CreatedAt,
	}, nil
}

func (handler Handler) UpdatePlaylistHandler(ctx *gin.Context, req models.UpdatePlaylistFormRequest) (models.PlaylistResponse, error) {
	//Find the playlist
	playlist, err := handler.query.GetPlaylistById(ctx, req.Id)
	if err != nil {
		// may handle ErrNoRow if have more time
		return models.PlaylistResponse{}, err
	}

	updatedPlaylist := models.PlaylistSchema{
		Id:           req.Id,
		PlaylistName: req.PlaylistName,
		Description:  req.Description,
		CoverImage:   playlist.CoverImage,
		IsPublic:     req.IsPublic,
	}

	//CASE: change audio file
	if req.CoverImage != nil {
		//remove previous file
		if playlist.CoverImage != "" {
			re := regexp.MustCompile(`(upload/images/.*)`)
			matches := re.FindString(playlist.CoverImage)

			if matches != "" {
				err := os.Remove(matches)
				if err != nil && !os.IsNotExist(err) {
					return models.PlaylistResponse{}, err
				}
			} else {
				return models.PlaylistResponse{}, err
			}
		}

		//Turn binaryfile into osFile and save to upload Folder
		filePath, err := utils.SaveUploadedFile(req.CoverImage, utils.UPLOAD_DIR_IMAGE)
		if err != nil {
			return models.PlaylistResponse{}, err
		}
		updatedPlaylist.CoverImage = filePath
	}

	res, err := handler.query.UpdatePlaylistQuery(ctx, updatedPlaylist)
	if err != nil {
		return models.PlaylistResponse{}, err
	}
	return models.PlaylistResponse{
		Id:           res.Id,
		PlaylistName: res.PlaylistName,
		Description:  res.Description,
		CoverImage:   res.CoverImage,
		IsPublic:     res.IsPublic,
		UserId:       res.UserId,
		CreatedAt:    res.CreatedAt,
	}, nil
}

func (handler Handler) DeletePlaylistByIdHandler(ctx *gin.Context, id int64) (models.DeletedResponse, error) {
	err := handler.query.DeletePlaylistByIdQuery(ctx, id)

	if err != nil {
		return models.DeletedResponse{Message: "Failed to delete!"}, err
	}

	return models.DeletedResponse{
		Message: "Deleted successfully",
	}, nil
}
