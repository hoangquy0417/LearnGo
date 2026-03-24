package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"example.com/learnGo/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductV1(ctx *gin.Context) {
	search := ctx.Query("search")
	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err := utils.ValidationRequired("Search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.ValidationStringLength("Search", search, 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.ValidationRegex("Search", search, searchRegex); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "limit must be a positive number",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Search": search,
		"Limit":  limit,
	})

}

func (u *ProductHandler) GetProductByIdV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"Message":    "Get Product (v1)",
		"Product id": id,
	})

}

func (u *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if err := utils.ValidationRegex("Slug", slug, slugRegex); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Get Product (v1)",
		"slug":    slug,
	})

}

func (u *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Create Product (v1)",
	})

}

func (u *ProductHandler) PutProductV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"Message":    "Upadate Product (v1)",
		"Product id": id,
	})

}

func (u *ProductHandler) DeleteProductV1(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, gin.H{
		"Message":    "Delete Product (v1)",
		"Product id": id,
	})

}
