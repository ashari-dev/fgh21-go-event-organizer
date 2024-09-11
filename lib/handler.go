package lib

import (
	"event-organizer/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOK(c *gin.Context, msg string, data any, pageInfo any) {
	c.JSON(http.StatusOK, dtos.Respont{
		Success:  true,
		Message:  msg,
		PageInfo: pageInfo,
		Result:   data,
	})
}

func HandlerNotfound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, dtos.Respont{
		Success: false,
		Message: msg,
	})
}

func HandlerUnauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, dtos.Respont{
		Success: false,
		Message: msg,
	})
}

func HandlerBadReq(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, dtos.Respont{
		Success: false,
		Message: msg,
	})
}
func HandlerMaxFile(c *gin.Context, msg string) {
	c.JSON(http.StatusRequestEntityTooLarge, dtos.Respont{
		Success: false,
		Message: msg,
	})
}
