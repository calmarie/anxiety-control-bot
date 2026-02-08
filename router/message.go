package router

import (
	"fmt"
	"log"
	"trevoga-control/anxiety"
	"trevoga-control/menu"
	"trevoga-control/state"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	// Логируем в консоль: от кого сообщение и какой текст.
	fmt.Println()
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	fmt.Println()

	if anxiety.HandleAnxietyMessage(bot, update) {
		return // сообщение обработано SM
	}

	switch {

	case update.Message.Text == "/start":
		menu.Start(bot, update.Message.Chat.ID)
		state.StateMachine[update.Message.Chat.ID] = []state.State{}
		state.StateMachine[update.Message.Chat.ID] = append(state.StateMachine[update.Message.Chat.ID], "start_")

	case update.Message.Text == "В начало":
		state.StateMachine[update.Message.Chat.ID] = []state.State{}
		state.StateMachine[update.Message.Chat.ID] = append(state.StateMachine[update.Message.Chat.ID], "start_")
		menu.SendFuncs(bot, update.Message.Chat.ID, 0, false)

	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не рабочее че то")
		bot.Send(msg)
	}

}
