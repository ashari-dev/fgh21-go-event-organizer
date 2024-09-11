package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetAllLocations() ([]models.Location, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM locations`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	locations, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Location])
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func GetOneLocations(id int) (models.Location, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM locations WHERE id = $1`

	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Location{}, err
	}

	locations, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Location])
	if err != nil {
		return models.Location{}, err
	}

	return locations, nil
}
