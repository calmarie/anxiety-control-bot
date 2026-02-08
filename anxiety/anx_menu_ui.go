package anxiety

import (
	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AnxietyMenu(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {

	row1 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Сделать запись", "diary_write"))
	row2 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Посмотреть записи", "diary_look"))
	row3 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Статистика", "diary_stats"))

	keys := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, navigation.RowBack)

	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		"Выбирай то, что хочешь сейчас:",
	)

	editMsg.ReplyMarkup = &keys
	bot.Send(editMsg)

}
