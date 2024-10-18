package apis

import (
	"MLMW/BEGoGin/infras/apis/controllers"
	"MLMW/BEGoGin/infras/auth"
	"MLMW/BEGoGin/infras/middlewares"
	"MLMW/BEGoGin/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// TODO: Server struct gonna have
// 3 layouts of api: Controller > Handler > Query
// api routers
// env config
// token maker
type Server struct {
	// here 3 controllers
	controller controllers.Controller
	router     *gin.Engine
	config     utils.ENVConfig
	tokenMaker auth.TokenMaker
}

func InitServerConnection(config utils.ENVConfig, store *sqlx.DB) (Server, error) {
	tokenMaker, err := auth.NewTokenMaker(config.SecretCode)
	if err != nil {
		return Server{}, err
	}

	server := Server{
		controller: controllers.NewController(store, config.SecretCode, &tokenMaker),
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setUpRouters()
	return server, nil
}

func (server *Server) setUpRouters() {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	public := router.Group("/apis")

	private := router.Group("/apis").Use(middlewares.AuthMiddleware(server.tokenMaker))
	// Serve static files from the "upload" directory at the URL path "/upload"
	router.Static("/upload", "./upload")

	//PUBLIC
	public.POST("/user/createUser", server.controller.CreateUserController)
	public.POST("/user/login", server.controller.LoginUserController)

	public.GET("/track/getAllPublicTracks", server.controller.GetAllPublicTracksController)
	public.GET("/track/getTrackById/:id", server.controller.GetTrackByIdController)

	public.GET("/playlist/getAllPublicPlaylists", server.controller.GetAllPublicPlaylists)
	public.GET("/playlist/getPlaylistById/:id", server.controller.GetPlaylistByIdController)
	public.GET("/playlist/getFullPlaylistDetail/:id", server.controller.GetFullPlaylistDetailController)

	public.GET("/search", server.controller.SearchingController)

	//PRIVATE authenticated users
	private.GET("/track/getAllTracksByUserId", server.controller.GetAllTracksByUserIdController)
	private.POST("/track/createTrack", server.controller.CreateTrackController)
	private.PUT("/track/updateTrack", server.controller.UpdateTrackController)
	private.DELETE("/track/deleteTrack/:id", server.controller.DeleteTrackByIdController)

	private.POST("/trackplaylist/upsert", server.controller.UpsertTracksPlaylistsController)
	private.GET("/trackplaylist/getDataTrackPlaylistById/:id", server.controller.GetDataTrackPlaylistByIdController)

	private.GET("/playlist/getAllPlaylistsByUserId", server.controller.GetAllPlaylistsByUserIdController)
	private.POST("/playlist/createPlaylist", server.controller.CreatePlaylistController)
	private.PUT("/playlist/updatePlaylist", server.controller.UpdatePlaylistController)
	private.DELETE("/playlist/deletePlaylist/:id", server.controller.DeletePlaylistByIdController)

	server.router = router
}

func (server Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
