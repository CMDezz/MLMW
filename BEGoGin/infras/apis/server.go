package apis

import (
	"MLMW/BEGoGin/infras/apis/controllers"
	"MLMW/BEGoGin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// TODO: Server struct gonna have
// 3 layouts of api: Controller > Handler > Query
// api routers
// env config
type Server struct {
	// here 3 controllers
	controller controllers.Controller
	router     *gin.Engine
	config     utils.ENVConfig
}

func InitServerConnection(config utils.ENVConfig, store *sqlx.DB) (Server, error) {
	server := Server{
		controller: controllers.NewController(store),
		config:     config,
	}

	server.setUpRouters()
	return server, nil
}

func (server *Server) setUpRouters() {
	router := gin.Default()
	public := router.Group("/apis")
	private := router.Group("/apis")

	//PUBLIC
	public.GET("/example", server.controller.TestUserController)

	//PRIVATE authenticated users
	// pRouter := router.Group("/").Use(authMidldeware())
	private.GET("/examplePrivate", func(ctx *gin.Context) { fmt.Println("samplePrivate") })

	server.router = router
}

func (server Server) Start(address string) error {
	err := server.router.Run(address)
	return err
}
