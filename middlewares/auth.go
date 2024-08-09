package middlewares

import (
	"net/http"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/gin-gonic/gin"
)

func tokenFailed(c *gin.Context) {
	if e := recover(); e != nil {
		c.JSON(http.StatusUnauthorized, lib.Respont{
			Success: false,
			Message: "Unauthorized",
		})
		c.Abort()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer tokenFailed(c)
		token := c.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			c.Set("UserId", userId)
			c.Next()
		} else {
			panic("Error: token invalid")
		}
	}
}
