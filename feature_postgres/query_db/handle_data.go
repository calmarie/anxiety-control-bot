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

func HandleUpdateAnxRow(ctx context.Context, conn *pgx.Conn, model models.AnxietyModel) {
	level := 0
	cause := "_"
	detCause := "_"
	sqlQuery := `
	SELECT * FROM anx ORDER BY time DESC LIMIT 1
	`
	row, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	var pgStruct models.AnxietyModel
	pgStruct, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.AnxietyModel])

	if pgStruct.Level != 0 {
		level = pgStruct.Level
	} else if model.Level != 0 {
		level = model.Level
	}

	if pgStruct.ShortReason != "_" {
		cause = pgStruct.ShortReason
	} else if model.ShortReason != "_" {
		cause = model.ShortReason
	}

	if pgStruct.DetailedReason != "_" {
		detCause = pgStruct.DetailedReason
	} else if model.DetailedReason != "_" {
		detCause = model.DetailedReason
	}

	sqlQuery = `
	UPDATE anx
	SET level = '$2'
	SET short_reason = '$3'
	SET detailed_reason = '$4'
	WHERE query_id = '$1'
	`

	_, err = conn.Exec(ctx, sqlQuery, pgStruct.QueryID, level, cause, detCause)

	if err != nil {
		panic(err)
	}
}
