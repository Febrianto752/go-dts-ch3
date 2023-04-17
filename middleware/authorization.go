package middleware

import (
	"net/http"
	"strconv"

	"github.com/Febrianto752/go-dts-ch3/config"
	"github.com/Febrianto752/go-dts-ch3/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.InitializeDB()

		productId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)

		if userData["role"] == "admin" {
			c.Next()
		}

		userID := uint(userData["id"].(float64))
		product := entity.Product{}

		err = db.Select("user_id").First(&product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}

		if product.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()

	}
}
