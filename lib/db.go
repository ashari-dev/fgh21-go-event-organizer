package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:@localhost:5432/event_organizer?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	return conn
}
