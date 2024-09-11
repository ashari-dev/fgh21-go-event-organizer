package routers

import (
	"event-organizer/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersNAtionality(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllNationality)

}
