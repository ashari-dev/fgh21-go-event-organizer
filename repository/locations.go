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

func CreateLocation(data models.Location) (models.Location, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO locations (name, image) VALUES ($1, $2) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Name, data.Image)
	if err != nil {
		return models.Location{}, err
	}

	location, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Location])
	if err != nil {
		return models.Location{}, err
	}

	return location, nil
}

func UpdateLocation(data models.Location, id int) (models.Location, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE locations SET (name, image) = ($1, $2) WHERE id = $3 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Name, data.Image, id)
	if err != nil {
		return models.Location{}, err
	}

	location, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Location])
	if err != nil {
		return models.Location{}, err
	}

	return location, nil
}

func DeleteLocation(id int) (models.Location, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM locations WHERE id = $1 RETURNING *`

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