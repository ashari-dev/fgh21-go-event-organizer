package routers

import "github.com/gin-gonic/gin"

func RouterCombain(r *gin.Engine) {
	RoutersUsers(r.Group("/users"))
	RoutersAuth(r.Group("/auth"))
	RoutersEvents(r.Group("/event"))
	RoutersProfile(r.Group("/profile"))
	RoutersNAtionality(r.Group("/nationality"))
	RoutersTransaction(r.Group("/transactions"))
	RoutersCategories(r.Group("/category"))
	RoutersLocations(r.Group("/locations"))
	RoutersWishlist(r.Group("/wishlist"))
}
