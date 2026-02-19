package querydb

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateAnxTable(ctx context.Context, conn *pgx.Conn) {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS anx (
    query_id SERIAL PRIMARY KEY,
    tg_id INTEGER NOT NULL,
    time TIMESTAMP NOT NULL,
    level INTEGER NOT NULL,
    short_reason VARCHAR NOT NULL,
    detailed_reason VARCHAR(1000)
	);
	`
	_, err := conn.Exec(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}

}
