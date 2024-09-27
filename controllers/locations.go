package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
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

func CreateLocations(c *gin.Context) {
	var data dtos.FormLocation
	err := c.Bind(&data)
	if err != nil {
		lib.HandlerBadReq(c, "bad request")
		return
	}

	location, err := repository.CreateLocation(models.Location{
		Name:  data.Name,
		Image: data.Image,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Failed to create location")
		return
	}
	lib.HandlerOK(c, "Create location successfully", location, nil)
}

func UpdateLocations(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data dtos.FormLocation
	err := c.Bind(&data)
	if err != nil {
		lib.HandlerBadReq(c, "bad request")
		return
	}

	location, err := repository.UpdateLocation(models.Location{
		Name:  data.Name,
		Image: data.Image,
	}, id)
	if err != nil {
		lib.HandlerBadReq(c, "Failed to update location")
		return
	}
	lib.HandlerOK(c, "Update location successfully", location, nil)
}

func DeleteLocations(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	location, err := repository.DeleteLocation(id)
	if err != nil {
		lib.HandlerNotfound(c, "location not found")
		return
	}

	lib.HandlerOK(c, "Delete location successfully", location, nil)
}
