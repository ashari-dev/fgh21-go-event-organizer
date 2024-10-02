package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func GetAllEvent(c *gin.Context) {
// 	events, err := repository.GetAllEvent()
// 	if err != nil {
// 		fmt.Println(err)
// 		lib.HandlerBadReq(c, "ini error")
// 		return
// 	}
// 	lib.HandlerOK(c, "Get all event", events, nil)
// }

func SearchEvents(c *gin.Context) {
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}

	events, count, err := repository.SearchEvents(search, limit, page)
	countTotalPage := math.Ceil(float64(count) / float64(limit))
	prev := page - 1
	next := page + 1
	if next > int(countTotalPage) {
		next = 0
	}
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "No events")
		return
	}

	lib.HandlerOK(c, "Search events", events, dtos.PageInfo{
		TotalData: count,
		TotalPage: int(countTotalPage),
		Page:      page,
		Limit:     limit,
		Next:      &next,
		Prev:      &prev,
	})
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
		lib.HandlerNotfound(c, "Event not found as")
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
