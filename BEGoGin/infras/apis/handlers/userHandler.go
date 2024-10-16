package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (handler Handler) TestUserHandler(ctx *gin.Context) {
	fmt.Println("TestUserHandler")

	handler.query.TestUserQuery(ctx)

}
