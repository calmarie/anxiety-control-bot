package anxiety

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleAnxityMenuCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	callbackQuery := update.CallbackQuery
	chatID := callbackQuery.Message.Chat.ID
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.Request(callback)
	suffix := strings.TrimPrefix(callbackQuery.Data, "diary_")
	switch {
	case suffix == "write":
		AskAnxietyLevel(bot, callbackQuery)
	case suffix == "look":
		SendAnxietyHistory(AnxietyStorage, bot, callbackQuery)
	case suffix == "stats":
		AnxietyStats(AnxietyStorage, chatID, bot)
	}

}

func HandleAnxietyMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) bool {
	chatID := update.Message.Chat.ID
	switch userStates[chatID] {
	case StateWaitingDetailedThought:

		SaveAnxietyData(
			chatID,
			tempLevel[chatID],
			tempCause[chatID],
			update.Message.Text,
		)

		SaveAnxietyToJSON(&AnxietyStorage)

		//для удобства отладки
		fmt.Println()
		fmt.Println()
		fmt.Println(chatID, tempLevel[chatID], tempCause[chatID], update.Message.Text)
		fmt.Println()
		//

		userStates[chatID] = StateIdle

		AskAnxietyFinish(bot, chatID)

		return true
	}
	return false

}

func HandleAnxietyCallback(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

	callbackQuery := update.CallbackQuery
	chatID := callbackQuery.Message.Chat.ID
	// Подтверждаем отклик
	callback := tgbotapi.NewCallback(callbackQuery.ID, "")
	bot.Request(callback)

	switch {

	case strings.HasPrefix(callbackQuery.Data, "anx_"):

		AskAnxietyCause(bot, callbackQuery)
		level, _ := strconv.Atoi(string(callbackQuery.Data[4]))
		tempLevel[chatID] = level
		userStates[chatID] = StateWaitingCauseCategory

	case strings.HasPrefix(callbackQuery.Data, "anxCause_"):

		cause := strings.TrimPrefix(callbackQuery.Data, "anxCause_")
		tempCause[chatID] = cause
		userStates[chatID] = StateWaitingDetailedThought

		AskAnxietyDetailedCause(bot, callbackQuery)

	default:
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "нет таких откликов"))
	}

}
