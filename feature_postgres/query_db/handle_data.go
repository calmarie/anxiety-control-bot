package querydb

import (
	"context"
	"trevoga-control/feature_postgres/models"

	"github.com/jackc/pgx/v5"
)

func HandleAnxRow(ctx context.Context, conn *pgx.Conn, model models.AnxietyModel) {
	sqlQuery := `
	INSERT INTO anx (tg_id, time, level, short_reason, detailed_reason)
    VALUES ($1, $2, $3, $4, $5)
	`

	_, err := conn.Exec(ctx, sqlQuery, model.ChatID, model.Time, model.Level, model.ShortReason, model.DetailedReason)

	if err != nil {
		panic(err)
	}
}
