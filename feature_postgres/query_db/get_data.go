package querydb

import (
	"context"
	"trevoga-control/feature_postgres/models"

	"github.com/jackc/pgx/v5"
)

func GetAnxDataFromDB(ctx context.Context, conn *pgx.Conn, chatID int64) []models.AnxietyModel {
	sqlQuery := `
	SELECT tg_id, time, level, short_reason, detailed_reason
    FROM anx 
	WHERE tg_id = $1
    ORDER BY time ASC 
	`

	rows, err := conn.Query(ctx, sqlQuery, int(chatID))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var storage []models.AnxietyModel
	for rows.Next() {
		var model models.AnxietyModel
		// type model struct
		err := rows.Scan(
			&model.ChatID,
			&model.Time,
			&model.Level,
			&model.ShortReason,
			&model.DetailedReason,
		)
		if err != nil {
			panic(err)
		}

		storage = append(storage, model)

	}

	return storage
}
