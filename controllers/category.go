package controllers

import (
	"event-organizer/lib"
	"event-organizer/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {
	categories, err := repository.GetAllCategory()
	if err != nil {
		lib.HandlerBadReq(c, "Failed request")
		return
	}

	lib.HandlerOK(c, "get all categories", categories, nil)
}

func GetOneCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	categories, err := repository.GetOneCategory(id)
	if err != nil {
		lib.HandlerNotfound(c, "category not found")
		return
	}

	lib.HandlerOK(c, "get category", categories, nil)
}
