package routers

import (
	"event-organizer/controllers"
	"event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersWishlist(rg *gin.RouterGroup) {
	rg.Use(middlewares.AuthMiddleware())
	rg.POST("", controllers.CreateWishlist)
	rg.GET("", controllers.GetAllWishlist)
	rg.DELETE("/:id", controllers.DeleteWishlist)
}
