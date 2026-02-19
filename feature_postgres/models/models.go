package models

import "time"

type AnxietyModel struct {
	ChatID         int
	Time           time.Time
	Level          int
	ShortReason    string
	DetailedReason string
}
