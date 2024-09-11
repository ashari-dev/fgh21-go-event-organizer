package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetAllCategory() ([]models.Categories, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM categories`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Categories])
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetOneCategory(id int) (models.Categories, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM categories WHERE id = $1`

	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Categories{}, err
	}

	categories, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Categories])
	if err != nil {
		return models.Categories{}, err
	}

	return categories, nil
}
