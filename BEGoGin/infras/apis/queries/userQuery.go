package queries

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (query Query) TestUserQuery(ctx *gin.Context) {
	fmt.Println("TestUserHandler")

}
