package routers

import (
	"event-organizer/controllers"
	"event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersTransaction(rg *gin.RouterGroup) {
	rg.POST("", middlewares.AuthMiddleware(), controllers.CreatTransaction)
	rg.GET("", middlewares.AuthMiddleware(), controllers.GetAllTransaction)

}
