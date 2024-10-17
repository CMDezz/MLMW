package handlers

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

func mapTrackSchemaToTrackResponse(tracksSchema *[]models.TrackSchema) []models.TrackResponse {
	tracksResponse := make([]models.TrackResponse, len(*tracksSchema))
	for i, track := range *tracksSchema {
		tracksResponse[i] = models.TrackResponse{
			Id:          track.Id,
			Title:       track.Title,
			Artist:      track.Artist,
			Album:       track.Album,
			Genre:       track.Genre,
			ReleaseYear: track.ReleaseYear,
			IsPublic:    track.IsPublic,
			UserId:      track.UserId,
			Url:         track.Url,
			CreatedAt:   track.CreatedAt,
		}
	}
	return tracksResponse
}

func (handler Handler) GetAllPublicTrackHandler(ctx *gin.Context) (models.GetAllPublicTracksResponse, error) {
	res, err := handler.query.GetAllPublicTrackQuery(ctx)

	if err != nil {
		return models.GetAllPublicTracksResponse{}, err
	}

	return models.GetAllPublicTracksResponse{
		Tracks: mapTrackSchemaToTrackResponse(&res),
	}, nil
}

func (handler Handler) GetAllTracksByUserIdHandler(ctx *gin.Context, id int64) (models.GetAllTracksResponse, error) {
	res, err := handler.query.GetAllTracksByUserIdQuery(ctx, id)

	if err != nil {
		return models.GetAllTracksResponse{}, err
	}

	return models.GetAllTracksResponse{
		Tracks: mapTrackSchemaToTrackResponse(&res),
	}, nil
}

func (handler Handler) DeleteTrackByIdHandler(ctx *gin.Context, id int64) (models.DeletedResponse, error) {
	err := handler.query.DeleteTrackByIdQuery(ctx, id)

	if err != nil {
		return models.DeletedResponse{Message: "Failed to delete!"}, err
	}

	return models.DeletedResponse{
		Message: "Deleted successfully",
	}, nil
}

func (handler Handler) CreateTrackHandler(ctx *gin.Context, req models.CreateTrackFormRequest) (models.TrackResponse, error) {

	//Turn binaryfile into osFile and save to upload Folder
	filePath, err := utils.SaveUploadedFile(req.TrackFile, utils.UPLOAD_DIR_TRACK)
	if err != nil {
		return models.TrackResponse{}, err
	}

	track := models.TrackSchema{
		Title:       req.Title,
		Artist:      req.Artist,
		Album:       req.Album,
		Genre:       req.Genre,
		ReleaseYear: req.ReleaseYear,
		UserId:      req.UserId,
		Url:         filePath,
	}

	res, err := handler.query.CreateTrackQuery(ctx, track)
	if err != nil {
		return models.TrackResponse{}, err
	}
	return models.TrackResponse{
		Id:          res.Id,
		Title:       res.Title,
		Artist:      res.Artist,
		Album:       res.Album,
		Genre:       res.Genre,
		IsPublic:    res.IsPublic,
		ReleaseYear: res.ReleaseYear,
		UserId:      res.UserId,
		Url:         res.Url,
		CreatedAt:   res.CreatedAt,
	}, nil
}

func (handler Handler) UpdateTrackHandler(ctx *gin.Context, req models.UpdateTrackFormRequest) (models.TrackResponse, error) {
	//Find the track
	track, err := handler.query.GetTrackById(ctx, req.Id)
	if err != nil {
		// may handle ErrNoRow if have more time
		return models.TrackResponse{}, err
	}

	updatedTrack := models.TrackSchema{
		Id:          req.Id,
		Title:       req.Title,
		Artist:      req.Artist,
		Album:       req.Album,
		Genre:       req.Genre,
		ReleaseYear: req.ReleaseYear,
		IsPublic:    req.IsPublic,
		Url:         track.Url,
	}

	//CASE: change audio file
	if req.TrackFile != nil {
		//remove previous file
		if track.Url != "" {
			re := regexp.MustCompile(`(upload/tracks/.*)`)
			matches := re.FindString(track.Url)

			if matches != "" {
				err := os.Remove(matches)
				if err != nil && !os.IsNotExist(err) { // no exist ? -> continue to upload, no problem
					return models.TrackResponse{}, err
				}
			} else {
				return models.TrackResponse{}, err
			}
		}

		//Turn binaryfile into osFile and save to upload Folder
		filePath, err := utils.SaveUploadedFile(req.TrackFile, utils.UPLOAD_DIR_TRACK)
		if err != nil {
			return models.TrackResponse{}, err
		}
		updatedTrack.Url = filePath
	}

	res, err := handler.query.UpdateTrackQuery(ctx, updatedTrack)
	if err != nil {
		return models.TrackResponse{}, err
	}
	return models.TrackResponse{
		Id:          res.Id,
		Title:       res.Title,
		Artist:      res.Artist,
		Album:       res.Album,
		Genre:       res.Genre,
		IsPublic:    res.IsPublic,
		ReleaseYear: res.ReleaseYear,
		UserId:      res.UserId,
		Url:         res.Url,
		CreatedAt:   res.CreatedAt,
	}, nil
}
