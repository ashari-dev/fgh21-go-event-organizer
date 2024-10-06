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

func InsertSection(data models.Section) (models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO event_sections (name, price, quantity, event_id) VALUES ($1, $2, $3, $4) RETURNING *`
	rows, err := db.Query(context.Background(), sql, data.Name, data.Price, data.Quantity, data.EventId)
	if err != nil {
		return models.Section{}, err
	}

	section, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Section])
	if err != nil {
		return models.Section{}, err
	}

	return section, nil
}

func UpdateSection(data models.Section, id int) (models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE event_sections SET (name, price, quantity, event_id) = ($1, $2, $3, $4) WHERE id = $5 RETURNING *`
	rows, err := db.Query(context.Background(), sql, data.Name, data.Price, data.Quantity, data.EventId, id)
	if err != nil {
		return models.Section{}, err
	}

	section, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Section])
	if err != nil {
		return models.Section{}, err
	}

	return section, nil
}

func DeleteSectionById(id int) (models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM event_sections WHERE id = $1 RETURNING *`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Section{}, err
	}

	section, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Section])
	if err != nil {
		return models.Section{}, err
	}

	return section, nil
}

func DeleteSectionByEventId(id int) (models.Section, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM event_sections WHERE event_id = $1 RETURNING *`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Section{}, err
	}

	section, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Section])
	if err != nil {
		return models.Section{}, err
	}

	return section, nil
}
