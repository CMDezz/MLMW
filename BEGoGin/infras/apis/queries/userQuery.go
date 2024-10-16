package queries

import (
	"MLMW/BEGoGin/models"
	"MLMW/BEGoGin/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (query Query) CreateUserQuery(ctx *gin.Context, user models.UserSchema) (models.UserSchema, error) {
	sqlCmd := fmt.Sprintf("INSERT INTO %s (username, full_name, email, hashed_password) VALUES($1,$2,$3,$4) RETURNING *", utils.TABLE_USERS)

	res := models.UserSchema{}

	err := query.store.QueryRowxContext(ctx, sqlCmd, user.Username, user.FullName, user.Email, user.HashedPassword).StructScan(&res)

	if err != nil {
		return models.UserSchema{}, err
	}
	return res, nil
}

func (query Query) GetUserByUsernameQuery(ctx *gin.Context, username string) (models.UserSchema, error) {
	sqlCmd := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", utils.TABLE_USERS)

	res := models.UserSchema{}

	err := query.store.QueryRowxContext(ctx, sqlCmd, username).StructScan(&res)

	if err != nil {
		return models.UserSchema{}, err
	}
	return res, nil
}

func (query Query) TestUserQuery(ctx *gin.Context) {
	fmt.Println("TestUserHandler")

}
