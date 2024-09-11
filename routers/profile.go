package routers

import (
	"event-organizer/controllers"
	"event-organizer/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersProfile(rg *gin.RouterGroup) {
	rg.GET("", middlewares.AuthMiddleware(), controllers.GetProfileByUserLogin)
	rg.PATCH("", middlewares.AuthMiddleware(), controllers.UpdateProfileByUserLogin)
	rg.PATCH("/upload-image", middlewares.AuthMiddleware(), controllers.UploadImageProfile)

}
