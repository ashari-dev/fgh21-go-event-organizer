package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAllUser(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		lib.HandlerBadReq(c, "error")
		return
	}

	lib.HandlerOK(c, "List all users", users, nil)

}

func ListOneUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := repository.GetOneUsers(id)
	if err != nil {
		lib.HandlerNotfound(c, "User not found")
		return
	}

	lib.HandlerOK(c, "List one user by id", users, nil)

}

func CreateUser(c *gin.Context) {
	var form dtos.FormUser
	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "bad request")
		return
	}
	roleId := 2
	data, err := repository.CreateUser(models.Users{
		Email:    form.Email,
		Username: *form.Username,
		Password: form.Password,
		RoleId:   roleId,
	})
	if err != nil {
		lib.HandlerBadReq(c, "create fiel")
		return
	}
	lib.HandlerOK(c, "User created successfully", data, nil)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var form dtos.FormUser
	err := c.Bind(&form)
	if err != nil {
		lib.HandlerBadReq(c, "bad request")
		return
	}
	data, err := repository.UpdateUser(models.Users{
		Email:    form.Email,
		Username: *form.Username,
		Password: form.Password,
	}, id)
	if err != nil {
		lib.HandlerBadReq(c, "Data not found")
		return
	}
	lib.HandlerOK(c, "User update successfully", data, nil)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := repository.DeleteUsers(id)
	if err != nil {
		lib.HandlerNotfound(c, "User not found")
		return
	}

	lib.HandlerOK(c, "User delete success", users, nil)

}

func ChangePassword(c *gin.Context) {
	id := c.GetInt("userId")
	users, err := repository.GetOneUsers(id)
	if err != nil {
		lib.HandlerNotfound(c, "User not found")
		return
	}

	var form dtos.FormChangePassword
	if err := c.Bind(&form); err != nil {
		lib.HandlerBadReq(c, "confirm password is wrong")
		return
	}

	isVerified := lib.Verify(form.OldPassword, users.Password)

	if !isVerified {
		lib.HandlerBadReq(c, "Old password is wrong")
		return
	} else {
		result, _ := repository.UpdateUserPassword(models.Users{
			Password: form.NewPassword,
		}, id)
		lib.HandlerOK(c, "Change user password is success", result, nil)
	}
}
