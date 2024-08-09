package models

import (
	"context"
	"fmt"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
)

type Profile struct {
	Id            int    `json:"id" db:"id"`
	Picture       *string `json:"picture" db:"picture"`
	FullName      string `json:"full-name" db:"full_name"`
	BirthDate     *string `json:"birth-day" db:"birth_date"`
	Gender        int    `json:"gender" db:"gender"`
	PhoneNumber   *string `json:"phone-number" db:"phone_number"`
	Profession    *string `json:"profession" db:"profession"`
	NationalityId *int    `json:"nationality-id" db:"nationality_id"`
	UsersId       int    `json:"users-id" db:"users_id"`
}

type JoinUserProfile struct{
	Id int 	`json:"id"`
	Email string `json:"email"`
	FullName string `json:"fullname"`
}

func CreateProfile(data Profile) (JoinUserProfile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	fmt.Println(data)

	initSQL := `INSERT INTO "profile" (picture,full_name,birth_date,gender,phone_number,profession,nationality_id,users_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := db.Exec(context.Background(), initSQL, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UsersId)

	if err != nil {
		fmt.Println(err)
	}

	sql:= `SELECT u.id, u.email, p.full_name FROM profile p JOIN users u ON p.users_id = u.id`
	profile := db.QueryRow(context.Background(),sql)

	var result JoinUserProfile
	profile.Scan(
		&result.Id,
		&result.Email,
		&result.FullName,
	)
	return result, err
}

func FindProfileByUserId(id int) {
	db := lib.DB()
	defer db.Close(context.Background())
}
