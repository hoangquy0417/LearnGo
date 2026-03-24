package main

import (
	v1handler "example.com/learnGo/internal/v1/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUserV1)
			user.GET("/:id", userHandlerV1.GetUserByIdV1)
			user.GET("admin/:uuid", userHandlerV1.GetUserByUuidV1)
			user.POST("", userHandlerV1.PostUserV1)
			user.PUT("/:id", userHandlerV1.PutUserV1)
			user.DELETE("/:id", userHandlerV1.DeleteUserV1)
		}
		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductV1)
			product.GET("/:slug", productHandlerV1.GetProductBySlugV1)
			product.POST("", productHandlerV1.PostProductV1)
			product.PUT("/:id", productHandlerV1.PutProductV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductV1)

		}

		categories := v1.Group("/categories")
		{
			categoriesHandlerV1 := v1handler.NewCategoryHandler()
			categories.GET("/:category", categoriesHandlerV1.GetCategoryV1)

		}

	}
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
