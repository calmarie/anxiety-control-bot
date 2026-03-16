package router

import (
	"context"
	"fmt"
	"log"
	"trevoga-control/anxiety"
	"trevoga-control/menu"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update, ctx context.Context, conn *pgx.Conn) {
	// Логируем в консоль: от кого сообщение и какой текст.
	fmt.Println()
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	fmt.Println()

	if anxiety.HandleAnxietyWriteDetCause(bot, update, ctx, conn) {
		return // сообщение обработано
	}

	switch {

	case update.Message.Text == "/start":
		menu.Start(bot, update.Message.Chat.ID)
		state.StateHistory[update.Message.Chat.ID] = []state.State{}
		state.StateHistory[update.Message.Chat.ID] = append(state.StateHistory[update.Message.Chat.ID], "start_")

	case update.Message.Text == "В начало":
		state.StateHistory[update.Message.Chat.ID] = []state.State{}
		state.StateHistory[update.Message.Chat.ID] = append(state.StateHistory[update.Message.Chat.ID], "start_")
		state.StateHistory[update.Message.Chat.ID] = append(state.StateHistory[update.Message.Chat.ID], state.StateStart)
		menu.SendFuncs(bot, update, false)

	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не рабочее че то")
		bot.Send(msg)
	}

}
