package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/ashari-dev/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string
}

func AuthLogin(c *gin.Context) {
	login := models.User{}
	c.Bind(&login)

	found := models.FindOneUserByEmail(login.Email)

	if found == (models.User{}) {
		c.JSON(http.StatusUnauthorized, lib.Respont{
			Success: false,
			Message: "Wrong emal and password",
		})
		return
	}

	isVerified := lib.Verify(login.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserTokenById(found.Id)
		fmt.Println(JWToken)
		c.JSON(http.StatusOK, lib.Respont{
			Success: true,
			Message: "login success",
			Results: Token{
				JWToken,
			},
		})
	} else {
		c.JSON(http.StatusUnauthorized, lib.Respont{
			Success: false,
			Message: "Wrong emal and password",
		})
	}
}

type FormRegister struct {
	FullName        string `form:"fullName"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmPassword" binding:"eqfield=Password"`
}

func RegisterUser(c *gin.Context) {
	form := FormRegister{}
	var user = models.User{}
	var profile = models.Profile{}
	err := c.Bind(&form)

	if err != nil {
		fmt.Println(nil)
	}

	user.Email = form.Email
	user.Password = form.Password
	profile.FullName = form.FullName
	createUser, _ := models.CreateUser(user)

	userId := createUser.Id
	profile.UsersId = userId
	
	creatProfile, _ := models.CreateProfile(profile)

	c.JSON(http.StatusOK, lib.Respont{
		Success: true,
		Message: "Create success",
		Results: creatProfile,
	})

}
