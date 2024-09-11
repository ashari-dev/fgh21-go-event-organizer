package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateProfile(data models.Profile) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO profile (picture, full_name, birth_date, gender, phone_number, profession, nationality_id, users_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Picture, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, data.UsersId)
	if err != nil {
		return models.Profile{}, err
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Profile])
	if err != nil {
		return models.Profile{}, err
	}

	return profile, nil
}

func GetProfileByUserId(id int) (models.JoinUserProfile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT u.id,p.picture, p.full_name, u."username", u."email",
			p.phone_number, p.gender, p.profession, p.nationality_id, p.birth_date 
			FROM users u JOIN profile p ON u.id = p.users_id
			WHERE u.id = $1`

	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.JoinUserProfile{}, err
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.JoinUserProfile])
	if err != nil {
		return models.JoinUserProfile{}, err
	}

	return profile, nil
}

func UpdateProfile(data models.Profile, id int) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE profile SET(full_name, birth_date, gender, 
			phone_number, profession, nationality_id)=($1,$2,$3,$4,$5,$6)
			WHERE users_id = $7 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.FullName, data.BirthDate, data.Gender, data.PhoneNumber, data.Profession, data.NationalityId, id)
	if err != nil {
		return models.Profile{}, err
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Profile])
	if err != nil {
		fmt.Println(err)
		return models.Profile{}, err
	}

	return profile, nil
}

func UploadImageProfile(data models.Profile, id int) (models.Profile, error) {
	db := lib.DB()
	defer db.Close(context.Background())
	sql := `UPDATE profile SET picture=$1
	WHERE users_id = $2 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Picture, id)
	if err != nil {
		return models.Profile{}, err
	}

	profile, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Profile])
	if err != nil {
		fmt.Println(err)
		return models.Profile{}, err
	}

	return profile, nil
}
