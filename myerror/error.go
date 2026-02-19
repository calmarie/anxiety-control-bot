package myerror

import (
	"fmt"
	"log"
	// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ReturnError(err error) {
	fmt.Println()
	fmt.Println()
	log.Println(err)
	fmt.Println()
	fmt.Println()

	// text := tgbotapi.NewMessage(chatID, "ваш запрос положил прод, мы уже сообщили прогеру, что он обосрался")
	// bot.Send(text)
}
