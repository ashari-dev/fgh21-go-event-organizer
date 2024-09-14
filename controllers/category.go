package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
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

func CreateCategory(c *gin.Context) {
	var form dtos.FormCategory
	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid request")
		return
	}

	category, err := repository.CreateCAtegory(models.Categories{
		Name: form.Name,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Failed to create category")
		return
	}

	lib.HandlerOK(c, "Create category success", category, nil)
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var form dtos.FormCategory
	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "Invalid request")
		return
	}

	category, err := repository.UpdateCategory(models.Categories{
		Name: form.Name,
	}, id)
	if err != nil {
		lib.HandlerBadReq(c, "Failed to update category")
		return
	}

	lib.HandlerOK(c, "Update category success", category, nil)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	categories, err := repository.DeleteCategory(id)
	if err != nil {
		lib.HandlerNotfound(c, "category not found")
		return
	}

	lib.HandlerOK(c, "Delete category success", categories, nil)
}
