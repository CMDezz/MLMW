package handlers

import (
	"MLMW/BEGoGin/infras/auth"
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (handler Handler) CreateUserHandler(ctx *gin.Context, req models.CreateUserRequest) (models.CreateUserResponse, error) {

	//Hash password when create an user
	hashed_password, err := auth.HashPassword(req.Password)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	newUser := models.UserSchema{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashed_password,
	}

	//query
	res, err := handler.query.CreateUserQuery(ctx, newUser)

	if err != nil {
		return models.CreateUserResponse{}, err
	}

	return models.CreateUserResponse{
		Username: res.Username,
		FullName: res.FullName,
		Email:    res.Email,
	}, nil

}

func (handler Handler) LoginUserHandler(ctx *gin.Context, req models.LoginUserRequest) (models.LoginUserResponse, error) {

	//find user
	user, err := handler.query.GetUserByUsernameQuery(ctx, req.Username)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	//compare password
	err = auth.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	//create new token
	accessToken, accessPayload, err := handler.tokenMaker.NewToken(user.Username, user.Email, utils.DEFAULT_ACCESS_TOKEN_DURATION)
	if err != nil {
		return models.LoginUserResponse{}, err
	}
	fmt.Println("--> ", accessToken)
	fmt.Println("--> ", accessPayload)
	fmt.Println("--> ", err)
	return models.LoginUserResponse{
		Token:     accessToken,
		ExpiredAt: accessPayload.ExpiredAt,
		Username:  accessPayload.Username,
		FullName:  user.FullName,
		Email:     user.Email,
	}, nil
}

func (handler Handler) TestUserHandler(ctx *gin.Context) {
	fmt.Println("TestUserHandler")

	handler.query.TestUserQuery(ctx)

}
