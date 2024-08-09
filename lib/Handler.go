package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOk(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, Respont{
		Success: true,
		Message: msg,
		Results: data,
	})
}

func HandlerNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Respont{
		Success: false,
		Message: msg,
	})
}
