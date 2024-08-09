package controllers

import (
	"strconv"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/ashari-dev/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListEvent(c *gin.Context) {
	result := models.ShowAllUser()

	lib.HandlerOk(c, "List all events", result)
}

func DetailEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := models.ShowOneEventById(id)

	if result.Id == 0 {
		lib.HandlerNotFound(c, "Event not found")
		return
	}
	lib.HandlerOk(c, "Get one event by id", result)
}

func CreateEvent(c *gin.Context) {
	event := models.Event{}
	c.Bind(&event)
	tes, _ := c.Get("UserId")
	createby, _ := tes.(int)
	event.CreatedBy = &createby
	result := models.InsertEvent(event)

	lib.HandlerOk(c, "Creater event success", result)
}

func UpdateEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	eventEdit := models.Event{}
	c.Bind(&eventEdit)

	result := models.EditEvent(id, eventEdit)

	if result.Id == 0 {
		lib.HandlerNotFound(c, "Event not found")
		return
	}

	lib.HandlerOk(c, "Event edit success", result)
}

func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result := models.RemoveEvent(id)

	if result.Id == 0 {
		lib.HandlerNotFound(c, "Event not found")
		return
	}

	lib.HandlerOk(c, "Event delete success", result)
}
