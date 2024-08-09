package models

import (
	"context"
	"fmt"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username"`
	Password string `json:"-" form:"password" binding:"min=8"`
}

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

	rows, _ := db.Query(context.Background(), "SELECT * FROM users ORDER BY id ASC")

	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[User])

	if err != nil {
		fmt.Println(err)
	}

	return users
}

func FindOneUserById(id int) User {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT id, email, username, "password" FROM users WHERE id = $1`
	row := db.QueryRow(context.Background(), sql, id)

	var user = User{}
	row.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
	)

	return user
}

func FindOneUserByEmail(email string) User {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT id, email, username, "password" FROM users WHERE email = $1`
	row := db.QueryRow(context.Background(), sql, email)

	var user = User{}
	row.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Password,
	)

	return user
}

func CreateUser(data User) (User, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	initSQL := `INSERT INTO users(email, username, "password") VALUES($1, $2, $3)`
	_, err := db.Exec(context.Background(), initSQL, data.Email, data.Username, data.Password)

	if err != nil {
		fmt.Println(err)
	}
	allUser := FindAllUsers()
	id := 0
	for _, v := range allUser {
		id = v.Id
	}
	data.Id = id
	return data, err
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
