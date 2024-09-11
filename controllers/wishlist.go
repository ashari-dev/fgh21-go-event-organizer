package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateWishlist(c *gin.Context) {
	var form dtos.FormWishlist
	c.Bind(&form)
	userId := c.GetInt("userId")

	allWislist, _ := repository.GetAllWishlist()
	var result []models.Wishlist
	for _, v := range allWislist {
		if v.UserId == userId {
			result = append(result, v)
		}
	}
	for _, v := range result {
		if v.EventId == form.EventId {
			lib.HandlerBadReq(c, "event already exists")
			return
		}
	}

	wishlist, err := repository.CreateWishlist(models.Wishlist{
		EventId: form.EventId,
		UserId:  userId,
	})
	if err != nil {
		lib.HandlerBadReq(c, "created failed")
		return
	}

	lib.HandlerOK(c, "Create successfuly", wishlist, nil)
}

func GetAllWishlist(c *gin.Context) {
	userId := c.GetInt("userId")
	wislist, err := repository.GetAllWishlistByUserLogin(userId)
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "get all wishlist failed")
		return
	}
	lib.HandlerOK(c, "get all wishlist successfuly", wislist, nil)
}
func DeleteWishlist(c *gin.Context) {
	// userId := c.GetInt("userId")
	id, _ := strconv.Atoi(c.Param("id"))

	wislist, err := repository.DeleteWishlist(id)
	if err != nil {
		fmt.Println(err)
		lib.HandlerNotfound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Delete wishlist success", wislist, nil)
}
