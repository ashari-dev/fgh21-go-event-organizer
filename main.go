package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Respont struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Results interface{} `json:"results,omitempty"`
}

type User struct{
	Id int `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"-" form:"password"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	data := []User{
		{Id: 1, Name: "Admin", Email: "admin@mail.com", Password: "1234"},
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, Respont{
			Success: true,
			Message: "OK",
		})
	})

	r.GET("/users",func(c *gin.Context) {
		c.JSON(http.StatusOK, Respont{
			Success: true,
			Message: "List data user",
			Results: data,
		})
	})

	r.GET("/users/:id",func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		cons := false
		index := -1
		for idx, item := range data{
			if id == item.Id {
				cons = true
				index = idx
			}
		}

		if cons {
			c.JSON(http.StatusOK, Respont{
				Success: true,
				Message: "User by id",
				Results: data[index],
			})
		}else{
			c.JSON(http.StatusNotFound, Respont{
				Success: false,
				Message: "Id is not found",
			})
		}
	})

	r.POST("/users",func(c *gin.Context) {
		user := User{}
		err := c.Bind(&user)
		result := 0
		for _, v := range data {
			result = v.Id
		}
		user.Id = result + 1


		cont := true
		for _, v := range data {
			if v.Email == user.Email {
				cont = false
			}
		}
		
		if err != nil {
			c.JSON(http.StatusBadRequest, Respont{
				Success: false,
				Message: "name and username not requitment",
			})
			fmt.Println(err)
		}else{
			if cont {
				data = append(data, user)
				c.JSON(http.StatusOK, Respont{
					Success: true,
					Message: "Cearte user success",
					Results: user,
				})
			}else{
				c.JSON(http.StatusUnauthorized, Respont{
					Success: false,
					Message: "Email already exist",
				})
			}
			fmt.Println(err)

		}

		
	})

	r.PATCH("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		cons := false
		idx := -1
		for index, item := range data{
			if id == item.Id {
				cons = true
				idx = index
			}
		}
		dataEdit := User{}
		err := c.Bind(&dataEdit)
		if err != nil {
			c.JSON(http.StatusBadRequest, Respont{
				Success: false,
				Message: "name and username not requitment",
			})
		}else{
			if cons {
				data[idx].Name = dataEdit.Name
				data[idx].Email = dataEdit.Email
				data[idx].Password = dataEdit.Password
				c.JSON(http.StatusOK,Respont{
					Success: true,
					Message: "data is update",
					Results: data[idx],
				})
			}else{
				c.JSON(http.StatusNotFound, Respont{
					Success: false,
					Message: "Id is not found",
				})
			}
		}
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		index := -1
		for idx, item := range data {
			if item.Id == id {
				index = idx
			}
		}

		if index != -1 {
			user := data[index]
			data = append(data[:index], data[index+1:]...)
			c.JSON(http.StatusOK, Respont{
				Success: true,
				Message: "Data user is delete",
				Results: user,
			})
		}else{
			c.JSON(http.StatusNotFound, Respont{
				Success: false,
				Message: "Id is not found",
			})
		}
	})


	r.Run("localhost:8888")
}