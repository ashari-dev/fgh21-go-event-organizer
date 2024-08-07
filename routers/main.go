package routers

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	users := r.Group("/users")
	UsersRouter(users)
}
