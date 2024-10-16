package middlewares

import (
	"MLMW/BEGoGin/infras/auth"
	"MLMW/BEGoGin/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenMaker auth.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//get header authorization & check
		authorizationHeader := ctx.GetHeader(utils.AUTHORIZATION_HEADER_KEY)
		if len(authorizationHeader) == 0 {
			err := fmt.Errorf("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		//check if valid authorization type
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := fmt.Errorf("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		//check beaer type
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != utils.AUTHORIZATION_BEAER_TYPE {
			err := fmt.Errorf("invalid authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		//check if token is valid
		accessToken := fields[1]
		payload, err := tokenMaker.ValidToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}

		//set context
		ctx.Set(utils.AUTHORIZATION_PAYLOAD_KEY, payload)
		//foward ctx to next handler
		ctx.Next()
	}
}
