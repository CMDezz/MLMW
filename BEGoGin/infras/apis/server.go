package apis

import (
	"MLMW/BEGoGin/infras/apis/controllers"
	"MLMW/BEGoGin/infras/auth"
	"MLMW/BEGoGin/infras/middlewares"
	"MLMW/BEGoGin/utils"
	"fmt"

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
	public := router.Group("/apis")
	private := router.Group("/apis").Use(middlewares.AuthMiddleware(server.tokenMaker))
	// Serve static files from the "upload" directory at the URL path "/upload"
	router.Static("/upload", "./upload")

	//PUBLIC
	public.GET("/example", server.controller.TestUserController)
	public.POST("/user/createUser", server.controller.CreateUserController)
	public.POST("/user/login", server.controller.LoginUserController)

	public.GET("/track/getAllPublicTracks", server.controller.GetAllPublicTracksController)
	public.GET("/track/getAllTracksByUserId/:id", server.controller.GetAllTracksByUserIdController)
	public.POST("/track/createTrack", server.controller.CreateTrackController)
	public.PUT("/track/updateTrack", server.controller.UpdateTrackController)
	public.DELETE("/track/deleteTrack/:id", server.controller.DeleteTrackByIdController)

	//PRIVATE authenticated users
	// pRouter := router.Group("/").Use(authMidldeware())
	private.GET("/examplePrivate", func(ctx *gin.Context) { fmt.Println("samplePrivate") })

	server.router = router
}

func (server Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
