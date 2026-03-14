package state

type State string

const (
	StateStart State = "start_"

	//модуль меню
	StateAnxietyMenu State = "menu_anxiety_"
	StateBreatheMenu State = "menu_breathe_"
	StatePraiseMenu  State = "menu_praise_"
	StateStatsMenu   State = "menu_stats_"

	//модуль меню -> Модуль тревоги
	StateAnxDieryWrite   State = "anx_diary_write_"
	StateAnxDieryHistory State = "anx_diary_history_"
	StateAnxDieryStats   State = "anx_diary_stats_"

	//модуль тревоги -> модуль записи тревоги
	StateAnxWriteLevel State = "anx_lvl_"

	StateAnxWriteCause         State = "anx_cause_"
	StateAnxWriteDetailedCause State = "anx_detailedCause_"

	//назад
	StateBack State = "back_"
)

var StateHistory = make(map[int64][]State)

// StateAnxWriteLevel1         State = "anx_lvl_1"
// StateAnxWriteLevel2         State = "anx_lvl_2"
// StateAnxWriteLevel3         State = "anx_lvl_3"
// StateAnxWriteLevel4         State = "anx_lvl_4"
// StateAnxWriteLevel5         State = "anx_lvl_5"
// StateAnxWriteLevel6         State = "anx_lvl_6"
// StateAnxWriteLevel7         State = "anx_lvl_7"
// StateAnxWriteLevel8         State = "anx_lvl_8"
// StateAnxWriteLevel9         State = "anx_lvl_9"
// StateAnxWriteLevel10        State = "anx_lvl_10"
