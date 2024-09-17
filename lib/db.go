package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)



func DB() *pgx.Conn {
	host := `3.25.204.209`
	port := `5432`
	user := `postgres`
	pass := `123`
	db := `find_your_event`

	
	conn, err := pgx.Connect(context.Background(), "postgresql://"+user+":"+pass+"@"+host+":"+port+"/"+db+"?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	return conn
}
