package controllers

import (
	"event-organizer/lib"
	"event-organizer/repository"

	"github.com/gin-gonic/gin"
)

func GetAllNationality(c *gin.Context) {
	nationalities, err := repository.GetAllNationality()
	if err != nil {
		lib.HandlerBadReq(c, "error")
		return
	}

	lib.HandlerOK(c, "Get all nationalities", nationalities, nil)
}
