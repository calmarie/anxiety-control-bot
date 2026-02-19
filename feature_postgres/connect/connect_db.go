package connect

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateConnect(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:asdfgh@localhost:5432/postgres")
}
