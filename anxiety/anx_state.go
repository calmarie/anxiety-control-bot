package anxiety

type AnxietyState int

const (
	StateIdle AnxietyState = iota
	StateWaitingCauseCategory
	StateWaitingDetailedThought
)

var userStates = make(map[int64]AnxietyState)
var tempLevel = make(map[int64]int)
var tempCause = make(map[int64]string)
