package navigation

import (
	// "fmt"
	// "trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var RowBack []tgbotapi.InlineKeyboardButton = tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("Вернуться назад 🔙", "back_"),
)
