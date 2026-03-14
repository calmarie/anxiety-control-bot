package models

import "time"

type AnxietyModel struct {
	QueryID        int
	ChatID         int
	Time           time.Time
	Level          int
	ShortReason    string
	DetailedReason string
}
