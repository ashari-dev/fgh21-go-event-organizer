package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllEvent() ([]models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM events`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return []models.Event{}, err
	}
	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Event])
	if err != nil {
		return []models.Event{}, err
	}
	return events, err
}

func GetAllEventByCreated(created int) ([]models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM events WHERE created_by = $1`
	rows, err := db.Query(context.Background(), sql, created)
	if err != nil {
		return []models.Event{}, err
	}
	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Event])
	if err != nil {
		return []models.Event{}, err
	}
	return events, err
}

func GetOneEvent(id int) (models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM events WHERE id = $1`
	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Event{}, err
	}

	events, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Event])
	if err != nil {
		return models.Event{}, err
	}
	return events, err
}

func CreateEvent(data models.Event) (models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO events (image, title, "date", description, location_id, created_by) 
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	row, err := db.Query(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, data.LocationId, data.CreatedBy)
	if err != nil {
		return models.Event{}, err
	}

	events, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Event])
	if err != nil {
		fmt.Println("sini")
		return models.Event{}, err
	}
	return events, err
}

func UpdateEvent(data models.Event, id int) (models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE events SET (image, title, "date", description, location_id) = 
			($1, $2, $3, $4, $5) WHERE id = $6 RETURNING *`
	row, err := db.Query(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, data.LocationId, id)
	if err != nil {
		return models.Event{}, err
	}

	events, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Event])
	if err != nil {
		fmt.Println(err)
		return models.Event{}, err
	}
	return events, err
}

func DeleteEvent(id int) (models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM events WHERE id = $1 RETURNING *`
	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Event{}, err
	}

	events, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Event])
	if err != nil {
		return models.Event{}, err
	}
	return events, err
}
