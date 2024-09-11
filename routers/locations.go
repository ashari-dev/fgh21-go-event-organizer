package routers

import (
	"event-organizer/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersLocations(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllLocations)
	rg.GET("/:id", controllers.GetOneLocations)

}
