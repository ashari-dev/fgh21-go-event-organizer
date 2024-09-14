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

func GetAllEvent(c *gin.Context) {
	events, err := repository.GetAllEvent()
	if err != nil {
		lib.HandlerBadReq(c, "error")
		return
	}

	lib.HandlerOK(c, "Get all event", events, nil)
}
func GetAllEventByCreated(c *gin.Context) {
	created := c.GetInt("userId")
	events, err := repository.GetAllEventByCreated(created)
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "No events")
		return
	}

	lib.HandlerOK(c, "Get all events by created", events, nil)
}

func GetOneEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	events, err := repository.GetOneEvent(id)
	if err != nil {
		lib.HandlerNotfound(c, "Event not found")
		return
	}

	lib.HandlerOK(c, "Get one event", events, nil)
}

func GetSectionEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	section, err := repository.GetSectionByEvent(id)
	if err != nil {
		lib.HandlerNotfound(c, "Section not found")
		return
	}

	lib.HandlerOK(c, "List all section", section, nil)
}

func CreateEvent(c *gin.Context) {
	createBy := c.GetInt("userId")
	var form dtos.FormEvent
	err := c.Bind(&form)
	if err != nil {
		fmt.Println(err)
	}

	event, err := repository.CreateEvent(models.Event{
		Image:       form.Image,
		Title:       form.Title,
		Date:        form.Date,
		Description: form.Description,
		LocationId:  form.LocationId,
		CreatedBy:   &createBy,
	})
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "Create Failed")
		return
	}

	lib.HandlerOK(c, "Create event success", event, nil)
}

func UpdateEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	createBy := c.GetInt("userId")

	var form dtos.FormEvent
	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "edit request failed")
		return
	}

	event, err := repository.GetOneEvent(id)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}
	if *event.CreatedBy != createBy {
		lib.HandlerBadReq(c, "cannot update other people's data")
		return
	}

	eventDel, err := repository.UpdateEvent(models.Event{
		Image:       form.Image,
		Title:       form.Title,
		Date:        form.Date,
		Description: form.Description,
		LocationId:  form.LocationId,
	}, id)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Update data success", eventDel, nil)
}

func DeleteEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	createBy := c.GetInt("userId")

	event, err := repository.GetOneEvent(id)

	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	if *event.CreatedBy != createBy {
		lib.HandlerBadReq(c, "cannot delete other people's data")
		return
	}

	eventDel, err := repository.DeleteEvent(id)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Delete data success", eventDel, nil)
}

func GetEventByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	events, err := repository.GetEventByCategory(id)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Get event by category", events, nil)
}

func CreateEventCategory(c *gin.Context) {
	var data dtos.EventCategory
	err := c.Bind(&data)
	if err != nil {
		lib.HandlerBadReq(c, "create request failed")
		return
	}

	eventCategory, err := repository.CreateEventCategory(models.EventCategory{
		EventId:    data.EventId,
		CategoryId: data.CategoryId,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Create Failed")
		return
	}
	lib.HandlerOK(c, "create event category successfully ", eventCategory, nil)

}
