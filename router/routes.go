package router

import (
	"context"
	"log"
	"strings"
	"trevoga-control/anxiety"
	"trevoga-control/menu"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

type HandlerFunc func(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn)

var routes = map[state.State]HandlerFunc{

	//queries:

	state.StateAnxietyMenu: anxiety.AnxietyMenu,
	state.StateBreatheMenu: HandlePlug,
	state.StatePraiseMenu:  HandlePlug,
	state.StateStatsMenu:   HandlePlug,

	state.StateAnxDieryWrite:   anxiety.AskAnxietyLevel,
	state.StateAnxDieryHistory: anxiety.SendAnxietyHistory,
	state.StateAnxDieryStats:   anxiety.AnxietyStats,

	state.StateAnxWriteCause:         anxiety.HandleAnxietyWriteCallback,
	state.StateAnxWriteLevel:         anxiety.HandleAnxietyWriteCallback,
	state.StateAnxWriteDetailedCause: anxiety.HandleAnxietyWriteCallback,
}

func AppendState(update *tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	data := update.CallbackQuery.Data
	if data == "" {
		log.Panic("пустой колбэк, прогер даун")
	}
	idx := strings.LastIndex(data, "_")

	if state.StateHistory[chatID][len(state.StateHistory)-1] != state.State(data[:idx]+"_") {
		log.Println(data)
		state.StateHistory[chatID] = append(state.StateHistory[chatID], state.State(data[:idx]+"_"))
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	if update.CallbackQuery == nil {
		log.Println("error: CallbackQuery == nil in HandleCallback")
		return
	}

	chatID := update.CallbackQuery.Message.Chat.ID

	log.Println(state.StateHistory)

	currentState := state.StateHistory[chatID][len(state.StateHistory[chatID])-1]
	log.Println(currentState)
	for stateRouter, handler := range routes {
		switch {
		case currentState == state.StateBack:
			log.Println("state_back__________________________________")
			state.StateHistory[chatID] = state.StateHistory[update.CallbackQuery.Message.Chat.ID][:len(state.StateHistory[chatID])-2]
			HandleCallback(bot, update, ctx, conn)
			return
		case currentState == state.StateStart:
			menu.SendFuncs(bot, update, true)

		case currentState == stateRouter:
			log.Println("handler started")
			handler(bot, update, ctx, conn)

		default:
			log.Println("error: State not found")
		}

	}

}

func HandlePlug(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	chatID := update.CallbackQuery.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "К сожалению, пока эта функция не работает")
	bot.Send(msg)
}
