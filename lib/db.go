package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:123@103.93.58.89:54326/find_your_event?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
