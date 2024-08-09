package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/ashari-dev/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	result := models.FindAllUsers()

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "List all users data",
		Results: result,
	})
}

func DetailUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := models.FindOneUserById(id)

	if result.Id == 0 {
		lib.HandlerNotFound(c, "User is not found")
		return
	}

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Data by user id",
		Results: result,
	})
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, _ := models.CreateUser(user)

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "User data added successfully",
		Results: data,
	})
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := models.UpdateData(user, id)

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "User update success",
		Results: data,
	})

}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := models.RemoveData(id)

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Data delete",
		Results: result,
	})
}
