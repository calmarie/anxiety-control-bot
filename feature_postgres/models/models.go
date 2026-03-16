package models

import "time"

type AnxietyModel struct {
	QueryID        int       `db:"query_reason"`
	ChatID         int       `db:"tg_id"`
	Time           time.Time `db:"time"`
	Level          int       `db:"level"`
	ShortReason    string    `db:"short_reason"`
	DetailedReason string    `db:"detailed_reason"`
}
