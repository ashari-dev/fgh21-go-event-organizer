package routers

import (
	"event-organizer/controllers"
	"event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersUsers(rg *gin.RouterGroup) {
	rg.GET("", controllers.ListAllUser)
	rg.GET("/:id", controllers.ListOneUser)
	rg.POST("", controllers.CreateUser)
	rg.DELETE("/:id", controllers.DeleteUser)
	rg.PATCH("/:id", controllers.UpdateUser)
	rg.PATCH("/update-password", middlewares.AuthMiddleware(), controllers.ChangePassword)

}
