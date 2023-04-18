package route

import (
	"github.com/Febrianto752/go-dts-ch3/handler"
	"github.com/Febrianto752/go-dts-ch3/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler handler.UserHandler, productHandler handler.ProductHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.UserRegisterHandler)
	router.POST("/signin", userHandler.UserLoginHandler)

	product := router.Group("/products")
	{
		product.Use(middleware.Authentication())
		product.POST("", productHandler.CreateProductHandler)
		product.GET("", middleware.ProductsAuthorization(), productHandler.GetProductsHandler)
		product.GET("/:id", middleware.ProductAuthorization(), productHandler.GetProductHandler)
		product.PUT("/:id", middleware.ProductAuthorization(), productHandler.UpdateProductHandler)
		product.DELETE("/:id", middleware.ProductAuthorization(), productHandler.DeleteProductHandler)
	}

	return router
}
