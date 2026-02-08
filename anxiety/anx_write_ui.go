package anxiety

import (
	"fmt"
	"strconv"
	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AskAnxietyLevel(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {

	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "anx_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "anx_2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "anx_3"),
		tgbotapi.NewInlineKeyboardButtonData("4", "anx_4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "anx_5"),
	)

	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("6", "anx_6"),
		tgbotapi.NewInlineKeyboardButtonData("7", "anx_7"),
		tgbotapi.NewInlineKeyboardButtonData("8", "anx_8"),
		tgbotapi.NewInlineKeyboardButtonData("9", "anx_9"),
		tgbotapi.NewInlineKeyboardButtonData("10", "anx_10"),
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(row1, row2, navigation.RowBack)

	// buf = tgbotapi.NewCopyMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID)

	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		"Какой уроень тревоги сейчас?",
	)

	editMsg.ReplyMarkup = &keys

	bot.Send(editMsg)

}

func AskAnxietyCause(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Здоровье", "anxCause_health"),
		tgbotapi.NewInlineKeyboardButtonData("Деньги", "anxCause_money"),
	)
	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Окружение", "anxCause_people"),
		tgbotapi.NewInlineKeyboardButtonData("Партнер", "anxCause_pair"),
	)
	row3 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Работа", "anxCause_job"),
		tgbotapi.NewInlineKeyboardButtonData("Образование", "anxCause_edu"),
	)
	row4 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Текущие обстоятельства", "anxCause_currant"),
		tgbotapi.NewInlineKeyboardButtonData("Другое", "anxCause_other"),
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, navigation.RowBack)

	level, _ := strconv.Atoi(string(callbackQuery.Data[4]))
	newText := fmt.Sprintf("Ваш уровень тревоги: %d/10. \n В чем причина тревоги?", level)
	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		newText,
	)
	editMsg.ReplyMarkup = &keys
	bot.Send(editMsg)

}

func AskAnxietyDetailedCause(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {

	chatID := callbackQuery.Message.Chat.ID
	editMsg := tgbotapi.NewEditMessageText(
		chatID,
		callbackQuery.Message.MessageID,
		"Опишите подробнее вашу тревожную мысль",
	)

	keys := tgbotapi.NewInlineKeyboardMarkup(navigation.RowBack)
	editMsg.ReplyMarkup = &keys
	bot.Send(editMsg)

}

func AskAnxietyFinish(bot *tgbotapi.BotAPI, chatID int64) {
	keys := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("В начало", "start_"),
			tgbotapi.NewInlineKeyboardButtonData("Посмотреть записи", "history_"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, "Записал. Спасибо 🙏")
	msg.ReplyMarkup = keys
	bot.Send(msg)
}
