package routers

import (
	"github.com/ashari-dev/fgh21-go-event-organizer/controllers"
	"github.com/ashari-dev/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func UsersRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddleware())
	routerGroup.GET("", controllers.ListUsers)
	routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("", controllers.CreateUser)
	routerGroup.PATCH("/:id", controllers.UpdateUser)
	routerGroup.DELETE(":id", controllers.DeleteUser)
}
