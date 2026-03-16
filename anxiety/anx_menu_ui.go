package anxiety

import (
	"context"
	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func AnxietyMenu(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	callbackQuery := update.CallbackQuery
	row1 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Сделать запись", "anx_diary_write_"))
	row2 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Посмотреть записи", "anx_diary_history_"))
	row3 := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Статистика", "anx_diary_stats_"))

	keys := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, navigation.RowBack)

	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		"Выбирай то, что хочешь сейчас:",
	)

	editMsg.ReplyMarkup = &keys
	bot.Send(editMsg)

}
