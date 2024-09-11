package repository

import (
	"context"
	"event-organizer/lib"
	"event-organizer/models"

	"github.com/jackc/pgx/v5"
)

func GetAllWishlist() ([]models.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT * FROM wishlist`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Wishlist])
	if err != nil {
		return nil, err
	}
	return wishlists, nil
}

func GetAllWishlistByUserLogin(UserId int) ([]models.WishlistJoinEvent, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `SELECT w.id, e.title, e."date", e.location_id location, 
			e.description FROM wishlist w 
			JOIN events e ON w.event_id = e.id 
			WHERE w.user_id = $1`

	rows, err := db.Query(context.Background(), sql, UserId)
	if err != nil {
		return nil, err
	}
	wishlists, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.WishlistJoinEvent])
	if err != nil {
		return nil, err
	}

	return wishlists, nil
}

func CreateWishlist(data models.Wishlist) (models.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO  wishlist (user_id, event_id) VALUES ($1, $2) RETURNING *`

	row, err := db.Query(context.Background(), sql, data.UserId, data.EventId)
	if err != nil {
		return models.Wishlist{}, err
	}

	wishlist, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Wishlist])
	if err != nil {
		return models.Wishlist{}, err
	}

	return wishlist, nil
}

func DeleteWishlist(id int) (models.Wishlist, error) {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `DELETE FROM wishlist WHERE id = $1 RETURNING *`

	row, err := db.Query(context.Background(), sql, id)
	if err != nil {
		return models.Wishlist{}, err
	}

	wishlist, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Wishlist])
	if err != nil {
		return models.Wishlist{}, err
	}

	return wishlist, nil
}
