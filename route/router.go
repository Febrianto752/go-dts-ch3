package route

import (
	"github.com/Febrianto752/go-dts-ch3/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/signup", userHandler.UserRegisterHandler)
	router.POST("/signin", userHandler.UserLoginHandler)

	return router
}
