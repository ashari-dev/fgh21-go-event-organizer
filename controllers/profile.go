package controllers

import (
	"event-organizer/dtos"
	"event-organizer/lib"
	"event-organizer/models"
	"event-organizer/repository"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetProfileByUserLogin(c *gin.Context) {
	id := c.GetInt("userId")

	profile, err := repository.GetProfileByUserId(id)
	if err != nil {
		lib.HandlerNotfound(c, "Get profile fail")
		return
	}
	lib.HandlerOK(c, "get profile by user login", profile, nil)
}

func UpdateProfileByUserLogin(c *gin.Context) {
	id := c.GetInt("userId")
	var form dtos.FormEditProfile
	if err := c.Bind(&form); err != nil {
		lib.HandlerBadReq(c, "Request error")
		return
	}

	user, err := repository.UpdateUserProfile(models.Users{
		Username: form.Username,
		Email:    form.Email,
	}, id)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	profile, err := repository.UpdateProfile(models.Profile{
		FullName:      form.Fullname,
		PhoneNumber:   &form.PhoneNumber,
		Gender:        &form.Gender,
		Profession:    &form.Profession,
		NationalityId: &form.Nationality,
		BirthDate:     &form.BirthDate,
	}, user.Id)
	if err != nil {
		fmt.Println("sini")
		lib.HandlerNotfound(c, "data not found")
		return
	}

	respont, err := repository.GetProfileByUserId(profile.UsersId)
	if err != nil {
		lib.HandlerNotfound(c, "data not found")
		return
	}

	lib.HandlerOK(c, "Update profile Success", respont, nil)

}

func UploadImageProfile(c *gin.Context) {
	userId := c.GetInt("userId")

	maxFile := 500 * 1024
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))

	file, err := c.FormFile("profileImg")
	if err != nil {
		if err.Error() == "http: request body too large" {
			lib.HandlerMaxFile(c, "file size too large, max capacity 500 kb")
			return
		}
		lib.HandlerBadReq(c, "not file to upload")
		return
	}
	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		lib.HandlerBadReq(c, "extension file not validate")
		return
	}

	newFile := uuid.New().String() + fileExt

	dirUpload := "./img/profile/"

	if err := c.SaveUploadedFile(file, dirUpload+newFile); err != nil {
		lib.HandlerBadReq(c, "file not upload")
		return
	}

	linkImg := "http://localhost:8080/img/profile/" + newFile

	delImgBefore, _ := repository.GetProfileByUserId(userId)
	if delImgBefore.Picture != nil {
		fileDel := strings.Split(*delImgBefore.Picture, "8080")[1]
		os.Remove("." + fileDel)
	}

	profile, _ := repository.UploadImageProfile(models.Profile{Picture: &linkImg}, userId)

	lib.HandlerOK(c, "Upload image success", profile, nil)

}
