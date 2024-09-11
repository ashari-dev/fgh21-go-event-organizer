package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetSectionByEvent(id int) ([]models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM event_sections WHERE event_id = $1`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return []models.Section{}, err
	}

	section, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Section])
	if err != nil {
		return []models.Section{}, err
	}

	return section, nil
}
