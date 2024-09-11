package routers

import (
	"event-organizer/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersAuth(rg *gin.RouterGroup) {
	rg.POST("/login", controllers.Login)
	rg.POST("/register", controllers.Register)
}
