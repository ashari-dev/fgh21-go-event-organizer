package routers

import (
	"github.com/ashari-dev/fgh21-go-event-organizer/controllers"
	"github.com/ashari-dev/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRoute(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware())
	r.GET("", controllers.ListEvent)
	r.GET("/:id", controllers.DetailEvent)
	r.POST("", controllers.CreateEvent)
	r.PATCH("/:id", controllers.UpdateEvent)
	r.DELETE("/:id", controllers.DeleteEvent)
}
