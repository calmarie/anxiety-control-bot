package anxiety

import (
	"context"

	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func AskAnxietyLevel(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	callbackQuery := update.CallbackQuery
	row1 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "anx_lvl_1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "anx_lvl_2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "anx_lvl_3"),
		tgbotapi.NewInlineKeyboardButtonData("4", "anx_lvl_4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "anx_lvl_5"),
	)

	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("6", "anx_lvl_6"),
		tgbotapi.NewInlineKeyboardButtonData("7", "anx_lvl_7"),
		tgbotapi.NewInlineKeyboardButtonData("8", "anx_lvl_8"),
		tgbotapi.NewInlineKeyboardButtonData("9", "anx_lvl_9"),
		tgbotapi.NewInlineKeyboardButtonData("10", "anx_lvl_10"),
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
		tgbotapi.NewInlineKeyboardButtonData("Здоровье", "anx_cause_health"),
		tgbotapi.NewInlineKeyboardButtonData("Деньги", "anx_cause_money"),
	)
	row2 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Окружение", "anx_cause_people"),
		tgbotapi.NewInlineKeyboardButtonData("Партнер", "anx_cause_pair"),
	)
	row3 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Работа", "anx_cause_job"),
		tgbotapi.NewInlineKeyboardButtonData("Образование", "anx_cause_edu"),
	)
	row4 := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Текущие обстоятельства", "anx_cause_currant"),
		tgbotapi.NewInlineKeyboardButtonData("Другое", "anx_cause_other"),
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, navigation.RowBack)

	newText := "Что вызвало тревогу?"
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
