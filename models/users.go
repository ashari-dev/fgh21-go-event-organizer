package models

import (
	"context"
	"fmt"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"min=8"`
}

// var dataUser = []User{
// 	{Id: 1, Name: "Admin", Email: "admin@mail.com", Password: "1234"},
// }

// func GetAllUsers() []User {
// 	data := dataUser

// 	return data
// }

// func GetOneUser(id int) User {
// 	data := dataUser

// 	user := User{}
// 	for _, item := range data {
// 		if id == item.Id {
// 			user = item
// 		}
// 	}

// 	return user
// }

// func CreateUser(data User) User {
// 	id := 0
// 	for _, v := range dataUser {
// 		id = v.Id
// 	}

// 	data.Id = id + 1
// 	dataUser = append(dataUser, data)

// 	return data
// }

// func RemoveData(id int) User {
// 	index := -1
// 	userDelete := User{}
// 	for idx, item := range dataUser {
// 		if item.Id == id {
// 			index = idx
// 			userDelete = item
// 		}
// 	}
// 	if userDelete.Id != 0 {
// 		dataUser = append(dataUser[:index], dataUser[index+1:]...)
// 	}

// 	return userDelete
// }

// func EditData(data User, id int) User {

// 	idx := -1

// 	for index, item := range dataUser {
// 		if id == item.Id {
// 			idx = index
// 		}
// 	}

// 	if idx == 0 {
// 		dataUser[idx].Name = data.Name
// 		dataUser[idx].Email = data.Email
// 		dataUser[idx].Password = data.Password
// 		data.Id = dataUser[idx].Id
// 	}

// 	return data
// }

func FindIndex(id int) int {
	data := FindAllUsers()

	index := -1
	for i, row := range data {
		if row.Id == id {
			index = i
		}
	}

	return index
}

func FindAllUsers() []User {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), "SELECT * FROM users")

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	return users
}

func FindOneUserById(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())
	index := FindIndex(id)

	rows, _ := db.Query(context.Background(), "SELECT * FROM users")

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}
	user := users[index]

	return user
}

func CreateUser(data User) User {
	db := lib.DB()
	defer db.Close(context.Background())

	initSQL := `INSERT INTO users(email, username, "password") VALUES($1, $2, $3)`
	db.Exec(context.Background(), initSQL, data.Email, data.Username, data.Password)
	allUser := FindAllUsers()
	id := 0
	for _, v := range allUser {
		id = v.Id
	}
	data.Id = id
	return data
}

func UpdateData(data User, id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	initSQL := `UPDATE users SET(email, username, "password")=($1, $2, $3) WHERE id=$4`
	db.Exec(context.Background(), initSQL, data.Email, data.Username, data.Password, id)

	data.Id = id
	return data
}

func RemoveData(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	userDel := FindOneUserById(id)
	initSQL := `DELETE FROM users WHERE id=$1`
	db.Exec(context.Background(), initSQL, id)

	return userDel
}
