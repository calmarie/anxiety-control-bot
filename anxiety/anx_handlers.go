package anxiety

import (
	"context"
	"log"
	"strconv"
	"strings"
	querydb "trevoga-control/feature_postgres/query_db"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

// func HandleAnxityMenuCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
// 	callbackQuery := update.CallbackQuery
// 	chatID := callbackQuery.Message.Chat.ID
// 	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
// 	_, err := bot.Request(callback)
// 	if err != nil {
// 		log.Println("error: Request callback bad")

// 	}
// 	suffix := strings.TrimPrefix(string(state.StateHistory[chatID][len(state.StateHistory[chatID])-1]), "anx_diary_")
// 	switch {
// 	case suffix == "write_":
// 		AskAnxietyLevel(bot, callbackQuery)
// 	case suffix == "history_":
// 		SendAnxietyHistory(bot, callbackQuery, ctx, conn)
// 	case suffix == "stats_":
// 		AnxietyStats(chatID, bot, ctx, conn)
// 	}

// }

func HandleAnxietyWriteDetCause(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) bool {
	chatID := update.Message.Chat.ID
	switch {
	case len(state.StateHistory[chatID]) == 0:
		return false

	case state.StateAnxWriteCause == state.StateHistory[chatID][len(state.StateHistory[chatID])-1]:

		model := SaveAnxietyToDB(
			chatID,
			0,
			"gap",
			update.Message.Text,
		)
		querydb.HandleUpdateAnxRow(ctx, conn, model)

		//для удобства отладки

		//

		AskAnxietyFinish(bot, chatID)

		return true
	}
	return false

}

func HandleAnxietyWriteCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	if update.CallbackQuery == nil {
		log.Println("error: HandleAnxietyWriteCallback CallbackQuery == nil")
	}

	callbackQuery := update.CallbackQuery
	chatID := callbackQuery.Message.Chat.ID

	log.Println(callbackQuery, chatID, "HandleAnxietyWriteCallback")

	// Подтверждаем отклик
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")

	_, err := bot.Request(callback)
	if err != nil {
		log.Println("error: Request callback bad")
	}

	switch {

	case state.StateAnxWriteLevel == state.StateHistory[chatID][len(state.StateHistory[chatID])-1]:

		AskAnxietyCause(bot, callbackQuery)
		if callbackQuery.Data != string(state.StateBack) {
			level, _ := strconv.Atoi(string(callbackQuery.Data[len(callbackQuery.Data)-1]))
			model := SaveAnxietyToDB(
				chatID,
				level,
				"gap",
				"gap",
			)

			querydb.HandleAnxRow(ctx, conn, model)
		} else {
			model := SaveAnxietyToDB(
				chatID,
				0,
				"gap",
				"gap",
			)
			querydb.HandleUpdateAnxRow(ctx, conn, model)
		}

	case state.StateAnxWriteCause == state.StateHistory[chatID][len(state.StateHistory[chatID])-1]:
		if callbackQuery.Data != string(state.StateBack) {
			cause := strings.TrimPrefix(callbackQuery.Data, "anx_cause_")
			model := SaveAnxietyToDB(
				chatID,
				0,
				cause,
				"gap",
			)
			querydb.HandleUpdateAnxRow(ctx, conn, model)
		}

		AskAnxietyDetailedCause(bot, callbackQuery)

	default:
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "нет таких откликов"))
	}

}
