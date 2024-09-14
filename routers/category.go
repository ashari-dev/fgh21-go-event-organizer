package routers

import (
	"event-organizer/controllers"

	"github.com/gin-gonic/gin"
)

func RoutersCategories(rg *gin.RouterGroup) {
	rg.GET("", controllers.GetAllCategory)
	rg.GET("/:id", controllers.GetOneCategory)
	rg.POST("", controllers.CreateCategory)
	rg.PATCH("/:id", controllers.UpdateCategory)
	rg.DELETE("/:id", controllers.DeleteCategory)

}
