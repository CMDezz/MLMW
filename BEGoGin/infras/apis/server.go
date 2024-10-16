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

	//PUBLIC
	public.GET("/example", server.controller.TestUserController)
	public.POST("/user/createUser", server.controller.CreateUserController)
	public.POST("/user/login", server.controller.LoginUserController)

	//PRIVATE authenticated users
	// pRouter := router.Group("/").Use(authMidldeware())
	private.GET("/examplePrivate", func(ctx *gin.Context) { fmt.Println("samplePrivate") })

	server.router = router
}

func (server Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
