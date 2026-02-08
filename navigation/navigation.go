package navigation

import (
	"fmt"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var RowBack []tgbotapi.InlineKeyboardButton = tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("Вернуться назад 🔙", "back_"),
)

func MoveBack(chatID int64, stateMachine map[int64][]state.State) state.State {

	if len(stateMachine[chatID]) <= 1 {
		fmt.Print("\n\n error blat \n\n")
	}
	fmt.Printf("\n\n %s \n\n", stateMachine[chatID])
	stateMachine[chatID] = stateMachine[chatID][:len(stateMachine[chatID])-1]

	state := stateMachine[chatID][len(stateMachine[chatID])-1]

	return state

}
