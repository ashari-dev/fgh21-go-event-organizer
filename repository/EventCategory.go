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

	sql := `SELECT ec.id, e.image, e.title, e."date",e.description, e.location_id, e.created_by FROM categories c 
			JOIN event_categories ec ON ec.event_id = c.id
			JOIN events e ON ec.event_id = e.id
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
