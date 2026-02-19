package router

import (
	"context"
	"strings"
	"trevoga-control/anxiety"
	"trevoga-control/menu"
	"trevoga-control/navigation"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func HandleCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID
	buf := true
	switch {
	case strings.HasPrefix(data, "anx_"),
		strings.HasPrefix(data, "anxCause_"):
		anxiety.HandleAnxietyCallback(bot, update)
	case strings.HasPrefix(data, "work_"):
		menu.HandleFuncCallback(bot, update)
	case strings.HasPrefix(data, "diary_"):
		anxiety.HandleAnxityMenuCallback(bot, update, ctx, conn)
	case strings.HasPrefix(data, "back_"):
		update.CallbackQuery.Data = string(navigation.MoveBack(chatID, state.StateMachine))
		buf = false
		HandleCallback(bot, update, ctx, conn)
	case strings.HasPrefix(data, "start_"):
		menu.SendFuncs(bot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, true)
	default:
		bot.Send(tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Неизвестная кнопка 🤷‍♂️",
		))
	}

	// отработка стэйтмашины для кнопки назад
	if buf {
		state.StateMachine[chatID] = append(state.StateMachine[chatID], state.State(update.CallbackQuery.Data))
	} else {
		state.StateMachine[chatID] = state.StateMachine[chatID][:len(state.StateMachine[chatID])-1]
	}

}
