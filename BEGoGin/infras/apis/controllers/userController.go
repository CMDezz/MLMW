package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) TestUserController(ctx *gin.Context) {
	fmt.Println("TestUserController")
	log.Println("TestUserController")
	controller.handler.TestUserHandler(ctx)
	ctx.JSON(http.StatusOK, nil)
}
