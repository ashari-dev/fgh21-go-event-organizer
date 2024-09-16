package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:123@3.25.204.209:5432/find_your_event?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
