package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetEventByCategory(id int) ([]models.Event, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT ec.id, e.image, e.title, e."date",e.description, e.location_id, e.created_by FROM events e 
			JOIN event_categories ec ON ec.event_id = e.id
			JOIN categories c ON c.id = ec.category_id
			WHERE ec.category_id = $1`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}

	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Event])
	if err != nil {
		return nil, err
	}

	return events, nil
}

func CreateEventCategory(data models.EventCategory) (models.EventCategory, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO event_categories (event_id, category_id) VALUES ($1, $2) RETURNING *`
	rows, err := db.Query(context.Background(), sql, data.EventId, data.CategoryId)
	if err != nil {
		return models.EventCategory{}, err
	}

	eventCategory, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.EventCategory])
	if err != nil {
		return models.EventCategory{}, err
	}

	return eventCategory, nil
}

func DeleteEventCategory(EventId int) (models.EventCategory, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM event_categories WHERE event_id = $1 RETURNING *`
	row, err := db.Query(context.Background(), sql, EventId)
	if err != nil {
		return models.EventCategory{}, err
	}

	events, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.EventCategory])
	if err != nil {
		return models.EventCategory{}, err
	}
	return events, err
}
