package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"Message": "List all users (v1)"})

}

func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get User (v1)",
		"user id": id,
	})

}

func (u *UserHandler) PostUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Create User (v1)",
	})

}

func (u *UserHandler) PutUserV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Upadate User (v1)",
		"user id": id,
	})

}

func (u *UserHandler) DeleteUserV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, gin.H{
		"Message": "Delete User (v1)",
		"user id": id,
	})

}
