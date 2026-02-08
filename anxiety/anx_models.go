package anxiety

type AnxietyEntry struct {
	Time           string `json:"time"`
	Level          int    `json:"level"`
	ShortReason    string `json:"short_reason"`
	DetailedReason string `json:"detailed_reason"`
}
