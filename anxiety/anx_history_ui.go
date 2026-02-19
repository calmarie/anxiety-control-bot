package anxiety

import (
	"context"
	"fmt"
	myerror "trevoga-control/Myerror"
	querydb "trevoga-control/feature_postgres/query_db"
	"trevoga-control/navigation"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func SendAnxietyHistory(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, ctx context.Context, conn *pgx.Conn) {
	chatID := callbackQuery.Message.Chat.ID
	text := "Вот что ты переживал за все это время:"
	anxData := querydb.GetAnxDataFromDB(ctx, conn, chatID)

	for _, v := range anxData {
		time := string(v.Time.Format("02/01/2006 15:04"))
		shortReason := DecryptAnxietyShortReason(v.ShortReason)
		detailedReason := v.DetailedReason
		text += fmt.Sprintf("\n\nВремя: %s", time)
		text += fmt.Sprintf("\nУровень: %d", v.Level)
		text += fmt.Sprintf("\nКраткая причина: %s", shortReason)
		text += fmt.Sprintf("\nПолное описание: %s", detailedReason)
		text += "\n---------------------"

	}
	text += fmt.Sprintf("\nВсего записей: %d", len(anxData))

	editMsg := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		text,
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(navigation.RowBack)
	editMsg.ReplyMarkup = &keys
	_, err := bot.Send(editMsg)
	if err != nil {
		myerror.ReturnError(err)
	}
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
