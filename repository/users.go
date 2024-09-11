package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetAllUsers() ([]models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users`

	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetOneUsers(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE id = $1`

	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Users{}, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func CreateUser(data models.Users) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `INSERT INTO users (email, username, password, role_id) VALUES ($1, $2, $3, $4) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Email, data.Username, data.Password, data.RoleId)
	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func DeleteUsers(id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM users WHERE id = $1 RETURNING *`

	rows, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Users{}, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}

func UpdateUser(data models.Users, id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `UPDATE users SET (email, username, password) = ($1, $2, $3) WHERE id = $4 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Email, data.Username, data.Password, id)
	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func UpdateUserProfile(data models.Users, id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())


	sql := `UPDATE users SET (email, username) = ($1, $2) WHERE id = $3 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Email, data.Username, id)
	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func UpdateUserPassword(data models.Users, id int) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	data.Password = lib.Encrypt(data.Password)

	sql := `UPDATE users SET password = $1 WHERE id = $2 RETURNING *`

	row, err := db.Query(context.Background(), sql, data.Password, id)
	if err != nil {
		return models.Users{}, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func GetUserByEmail(email string) (models.Users, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM users WHERE email = $1`

	rows, err := db.Query(context.Background(), sql, email)
	if err != nil {
		return models.Users{}, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Users])
	if err != nil {
		return models.Users{}, err
	}

	return users, nil
}
