package controllers

import (
	"event-organizer/lib"
	"event-organizer/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllLocations(c *gin.Context) {
	locations, err := repository.GetAllLocations()
	if err != nil {
		lib.HandlerBadReq(c, "Request failed")
		return
	}

	lib.HandlerOK(c, "Get all locations successfully", locations, nil)
}

func GetOneLocations(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	locations, err := repository.GetOneLocations(id)
	if err != nil {
		lib.HandlerNotfound(c, "location not found")
		return
	}

	lib.HandlerOK(c, "Get one location successfully", locations, nil)
}
