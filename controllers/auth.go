package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var formLogin dtos.Login
	err := c.Bind(&formLogin)
	if err != nil {
		lib.HandlerBadReq(c, "bad request")
		return
	}

	found, _ := repository.GetUserByEmail(formLogin.Email)
	if found == (models.Users{}) {
		lib.HandlerUnauthorized(c, "Email is wrong")
		return
	}

	isVerified := lib.Verify(formLogin.Password, found.Password)
	if !isVerified {
		lib.HandlerUnauthorized(c, "Password is wrong")
		return
	} else {
		JWToken := lib.GenerateUserTokenById(found.Id)
		lib.HandlerOK(c, "Login success", dtos.Token{Token: JWToken}, nil)
	}
}

func Register(c *gin.Context) {
	var formReg dtos.Register
	err := c.Bind(&formReg)
	if err != nil {
		lib.HandlerBadReq(c, "Confirm password wrong")
		return
	}

	user, err := repository.CreateUser(models.Users{
		Email:    formReg.Email,
		Password: formReg.Password,
		RoleId:   2,
	})
	if err != nil {
		lib.HandlerBadReq(c, "Email already in use")
		return
	}

	profile, err := repository.CreateProfile(models.Profile{
		FullName: formReg.FullName,
		UsersId:  user.Id,
	})
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "Email already in use")
		return
	}

	data, err := repository.GetProfileByUserId(profile.UsersId)
	if err != nil {
		fmt.Println(err)
		lib.HandlerBadReq(c, "Email already in use")
		return
	}

	lib.HandlerOK(c, "Register success", data, nil)
}
