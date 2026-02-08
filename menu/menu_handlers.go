package menu

import (
	"trevoga-control/anxiety"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleFuncCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	chatID := callbackQuery.Message.Chat.ID
	// Подтверждаем отклик
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.Request(callback)
	msg := tgbotapi.NewMessage(chatID, "нажми на кнопочку пж")
	switch {
	case string(callbackQuery.Data) == "work_anxiety":
		anxiety.AnxietyMenu(bot, callbackQuery)
	case string(callbackQuery.Data) == "work_breathe":
		msg = tgbotapi.NewMessage(chatID, "К сожалению, пока эта функция не работает")
		bot.Send(msg)
	case string(callbackQuery.Data) == "work_praise":
		msg = tgbotapi.NewMessage(chatID, "К сожалению, пока эта функция не работает")
		bot.Send(msg)
	case string(callbackQuery.Data) == "work_stats":
		msg = tgbotapi.NewMessage(chatID, "К сожалению, пока эта функция не работает")
		bot.Send(msg)
	}
}
