package routers

import (
	"event-organizer/controllers"
	"event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersEvents(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllEvent)
	rg.GET("/:id", controllers.GetOneEvent)
	rg.GET("/category/:id", controllers.GetEventByCategory)
	rg.GET("/section/:id", controllers.GetSectionEvent)
	rg.POST("", middlewares.AuthMiddleware(), controllers.CreateEvent)
	rg.GET("/my-events", middlewares.AuthMiddleware(), controllers.GetAllEventByCreated)
	rg.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteEvent)
	rg.PATCH("/:id", middlewares.AuthMiddleware(), controllers.UpdateEvent)

}
