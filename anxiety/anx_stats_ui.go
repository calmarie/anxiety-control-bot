package anxiety

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AnxietyStats(anxietyStorage map[int64][]AnxietyEntry, chatID int64, bot *tgbotapi.BotAPI) {
	// Implement statistics calculation logic here
	text := "ntrn"
	if len(anxietyStorage) == 0 {
		text = "Нет данных для статистики."
	} else {
		max_lvl := 0
		min_lvl := 10
		popular_cause := ""
		// week_max_lvl := 0
		// week_min_lvl := 10
		// week_popular_cause := ""
		day_max_lvl := 0
		day_min_lvl := 10
		day_popular_cause := ""
		text = "Статистика по тревоге:\n"

		entries := anxietyStorage[chatID]
		causes := make(map[string]int)
		day_causes := make(map[string]int)
		for _, entry := range entries {

			if entry.Level > max_lvl {
				max_lvl = entry.Level
				if string(time.Now().Format("02/01/2006")) == strings.SplitN(entry.Time, " ", 2)[0] {
					day_max_lvl = entry.Level
				}
			}
			if entry.Level < min_lvl {
				min_lvl = entry.Level
				if string(time.Now().Format("02/01/2006")) == strings.SplitN(entry.Time, " ", 2)[0] {
					day_min_lvl = entry.Level
				}
			}
			if string(time.Now().Format("02/01/2006")) == strings.SplitN(entry.Time, " ", 2)[0] {
				day_causes[entry.ShortReason] += 1
			}
			causes[entry.ShortReason] += 1

		}

		for i, v := range causes {
			max_v := 0
			if v > max_v {
				popular_cause = i
			}
		}

		for i, v := range day_causes {
			max_v := 0
			if v > max_v {
				day_popular_cause = i
			}
		}
		text += "\n\nТревога за все время:"
		text += fmt.Sprintf("\nМаксимальный уровень: %s", strconv.Itoa(max_lvl))
		text += fmt.Sprintf("\nМинимальный уровень: %s", strconv.Itoa(min_lvl))
		text += fmt.Sprintf("\nСамая частая причина: %s", popular_cause)
		text += "\n\nТревога за день:"
		text += fmt.Sprintf("\nМаксимальный уровень: %s", strconv.Itoa(day_max_lvl))
		text += fmt.Sprintf("\nМинимальный уровень: %s", strconv.Itoa(day_min_lvl))
		text += fmt.Sprintf("\nСамая частая причина: %s", day_popular_cause)

		fmt.Println(causes)
		fmt.Println(day_causes)
	}

	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}
