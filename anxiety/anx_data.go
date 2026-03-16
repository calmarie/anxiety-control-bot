package anxiety

import (
	"time"
	"trevoga-control/feature_postgres/models"
)

func SaveAnxietyToDB(chatID int64, templevel int, tempcause string, longcause string) models.AnxietyModel {
	model := models.AnxietyModel{
		QueryID:        0,
		ChatID:         int(chatID),
		Time:           time.Now(),
		Level:          templevel,
		ShortReason:    tempcause,
		DetailedReason: longcause,
	}

	return model
}
