package v1handler

import (
	"net/http"

	"example.com/learnGo/utils"
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
	idStr := ctx.Param("id")

	id, err := utils.ValidationPositiveInt("ID", idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get User (v1)",
		"user id": id,
	})

}

func (u *UserHandler) GetUserByUuidV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")

	uid, err := utils.ValidationUuid("Uuid", uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get User (v1)",
		"user id": uid,
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
