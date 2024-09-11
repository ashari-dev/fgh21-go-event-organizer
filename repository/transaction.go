package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func CreateTransaction(data models.Transaction) (models.Transaction, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO transactions (event_id, payment_method_id, user_id) VALUES ($1, $2, $3) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.EventId, data.PaymentMethodId, data.UserId)
	if err != nil {
		return models.Transaction{}, err
	}

	transaction, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, err
}

func GetTransaction(id int) (models.TransactionJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
			SELECT t.id, p.full_name, e.title as "event_title", e.location_id, 
			e."date",pm.name as "payment_method", array_agg(es.name) as "section_name", 
			array_agg(td.ticket_qty) as "ticket_qty" 
			FROM transactions t
			JOIN users u on u.id = t.user_id
			JOIN profile p ON p.users_id = u.id
			JOIN events e ON e.id = t.event_id
			JOIN payment_method pm ON pm.id = t.payment_method_id
			JOIN transaction_details td ON td.transaction_id = t.id
			JOIN event_sections es ON es.id = td.section_id
			WHERE t.id = $1
			GROUP BY t.id, p.full_name, e.title, e.location_id, e."date",pm.name
			`
	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.TransactionJoin{}, err
	}

	transaction, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.TransactionJoin])
	if err != nil {
		return models.TransactionJoin{}, err
	}

	return transaction, nil
}

func GetAllTransaction(id int) ([]models.TransactionJoin, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `
			SELECT t.id, p.full_name, e.title as "event_title", e.location_id, 
			e."date",pm.name as "payment_method", array_agg(es.name) as "section_name", 
			array_agg(td.ticket_qty) as "ticket_qty" 
			FROM transactions t
			JOIN users u on u.id = t.user_id
			JOIN profile p ON p.users_id = u.id
			JOIN events e ON e.id = t.event_id
			JOIN payment_method pm ON pm.id = t.payment_method_id
			JOIN transaction_details td ON td.transaction_id = t.id
			JOIN event_sections es ON es.id = td.section_id
			WHERE u.id = $1
			GROUP BY t.id, p.full_name, e.title, e.location_id, e."date",pm.name
			`
	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return []models.TransactionJoin{}, err
	}

	transaction, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionJoin])
	if err != nil {
		return []models.TransactionJoin{}, err
	}

	return transaction, nil
}