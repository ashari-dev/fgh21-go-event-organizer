package main

import (
	"github.com/ashari-dev/fgh21-go-event-organizer/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	routers.RouterCombain(r)

	r.Run("localhost:8888")
}