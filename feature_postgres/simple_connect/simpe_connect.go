package simpleconnect

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnect() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:asdfgh@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("БД подключена успешно")
}
