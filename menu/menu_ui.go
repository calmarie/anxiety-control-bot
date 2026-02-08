package menu

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, chatID int64) {

	text := `🧠 Добро пожаловать в бот который поможет в проработке тревоги!

	📊 Этот бот поможет вам:
	✅ Отслеживать уровень тревоги
	✅ Анализировать причины тревожных мыслей
	✅ Вести дневник эмоционального состояния

	⚠️ ВАЖНО ЗНАТЬ:
	❌ Не заменяет психолога или врача
	❌ Не ставит диагнозы
	❌ Не даёт медицинских рекомендаций

	🔥 Если тревога выше 7-8 баллов из 10 — возможно лучше обратиться за поддержкой!

	💡 Бот = инструмент для осознанности, а не лечение.
	
	👇Жми "В начало" чтобы приступить👇`

	msg := tgbotapi.NewMessage(chatID, text)

	button1 := tgbotapi.NewKeyboardButton("В начало")
	button2 := tgbotapi.NewKeyboardButton("О боте")

	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(button1), tgbotapi.NewKeyboardButtonRow(button2))
	keyboard.ResizeKeyboard = true
	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}

func SendFuncs(bot *tgbotapi.BotAPI, chatID int64, msgID int, back bool) {
	row_anxiety := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Поработать с тревогой", "work_anxiety"),
	)
	row_breathe := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Заземлиться/подышать", "work_breathe"),
	)
	row_praise := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Дневник похвалы", "work_praise"),
	)
	row_stats := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Посмотреть статистику", "work_stats"),
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(row_anxiety, row_breathe, row_praise, row_stats)
	if back {
		editMsg := tgbotapi.NewEditMessageText(chatID, msgID, "Выбирай то, что хочешь сейчас:")
		editMsg.ReplyMarkup = &keys
		_, err := bot.Send(editMsg)
		if err != nil {
			log.Println(err)
		}
	} else {
		msg := tgbotapi.NewMessage(chatID, "Выбирай то, что хочешь сейчас:")
		msg.ReplyMarkup = keys
		bot.Send(msg)
	}

}
