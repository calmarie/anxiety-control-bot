package anxiety

import (
	"fmt"
	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendAnxietyHistory(anxietyStorage map[int64][]AnxietyEntry, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	chatID := callbackQuery.Message.Chat.ID
	text := "Вот что ты переживал за все это время:"
	AnxietyHistory := anxietyStorage[chatID]
	for _, v := range AnxietyHistory {
		time := v.Time
		shortReason := DecryptAnxietyShortReason(v.ShortReason)
		detailedReason := v.DetailedReason
		text += fmt.Sprintf("\n\nВремя: %s", time)
		text += fmt.Sprintf("\nУровень: %d", v.Level)
		text += fmt.Sprintf("\nКраткая причина: %s", shortReason)
		text += fmt.Sprintf("\nПолное описание: %s", detailedReason)
		text += "\n---------------------"

	}
	text += fmt.Sprintf("\nВсего записей: %d", len(AnxietyHistory))

	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		text,
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(navigation.RowBack)
	editMsg.ReplyMarkup = &keys
	bot.Send(editMsg)
}

func DecryptAnxietyShortReason(code string) string {
	switch code {
	case "health":
		return "Здоровье"
	case "money":
		return "Деньги"
	case "people":
		return "Окружение"
	case "pair":
		return "Партнер"
	case "job":
		return "Работа"
	case "edu":
		return "Образование"
	case "currant":
		return "Текущие обстоятельства"
	default:
		return "Прочее"
	}
}
