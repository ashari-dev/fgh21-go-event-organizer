package models

import (
	"context"
	"fmt"

	"github.com/ashari-dev/fgh21-go-event-organizer/lib"
	"github.com/jackc/pgx/v5"
)

type Event struct {
	Id          int    `json:"id" db:"id"`
	Image       string `json:"image" form:"image" db:"image"`
	Title       string `json:"title" form:"title" db:"title"`
	Date        string `json:"date" form:"date" db:"date"`
	Description string `json:"description" form:"description" db:"description"`
	LocationId  *int   `json:"location_id"  form:"location_id" db:"location_id"`
	CreatedBy   *int   `json:"create_by" form:"create_by" db:"create_by"`
}

func ShowAllUser() []Event {
	db := lib.DB()
	defer db.Close(context.Background())

	rows, _ := db.Query(context.Background(), "SELECT * FROM events ORDER BY id ASC")

	event, err := pgx.CollectRows(rows, pgx.RowToStructByPos[Event])

	if err != nil {
		fmt.Println(err)
	}

	return event
}

func ShowOneEventById(id int) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	allEvent := ShowAllUser()

	var event Event
	for _, v := range allEvent {
		if v.Id == id {
			event = v
		}
	}

	return event
}

func InsertEvent(event Event) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `INSERT INTO events("image", "title", "date", "description")
	VALUES ($1, $2, $3, $4)`
	db.Exec(context.Background(), sql, event.Image, event.Title, event.Date, event.Description)
	id := 0
	for _, v := range ShowAllUser() {
		id = v.Id
	}
	event.Id = id

	return event
}

func EditEvent(id int, data Event) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	sql := `UPDATE events SET ("image", "title", "date", "description") = ($1, $2, $3, $4) WHERE id=$5`

	db.Exec(context.Background(), sql, data.Image, data.Title, data.Date, data.Description, id)

	data.Id = id
	return data
}

func RemoveEvent(id int) Event {
	db := lib.DB()
	defer db.Close(context.Background())

	data := ShowOneEventById(id)

	sql := `DELETE FROM events WHERE id=$1`

	db.Exec(context.Background(), sql, id)

	return data
}
