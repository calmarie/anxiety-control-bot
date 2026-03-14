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

func SendFuncs(bot *tgbotapi.BotAPI, update *tgbotapi.Update, back bool) {

	row_anxiety := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Поработать с тревогой", "menu_anxiety_"),
	)
	row_breathe := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Заземлиться/подышать", "menu_breathe_"),
	)
	row_praise := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Дневник похвалы", "menu_praise_"),
	)
	row_stats := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Посмотреть статистику", "menu_stats_"),
	)
	keys := tgbotapi.NewInlineKeyboardMarkup(row_anxiety, row_breathe, row_praise, row_stats)

	if back {
		chatID := update.CallbackQuery.Message.Chat.ID
		msgID := update.CallbackQuery.Message.MessageID
		editMsg := tgbotapi.NewEditMessageText(chatID, msgID, "Выбирай то, что хочешь сейчас:")
		editMsg.ReplyMarkup = &keys
		_, err := bot.Send(editMsg)
		if err != nil {
			log.Println(err)
		}
	} else {
		chatID := update.Message.Chat.ID

		msg := tgbotapi.NewMessage(chatID, "Выбирай то, что хочешь сейчас:")
		msg.ReplyMarkup = keys
		bot.Send(msg)
	}

}
