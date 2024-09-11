package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetAllNationality() ([]models.Nationality, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM nationalities`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return []models.Nationality{}, err
	}

	nationalities, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Nationality])
	if err != nil {
		return []models.Nationality{}, err

	}
	return nationalities, nil

}
