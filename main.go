package main

import (
	"context"
	"log"
	"os"
	"time"
	"trevoga-control/feature_postgres/connect"
	querydb "trevoga-control/feature_postgres/query_db"
	"trevoga-control/router"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

// точка входа
func main() {

	//выгружаю данные из файла
	ctx := context.Background()
	conn, err := connect.CreateConnect(ctx)
	if err != nil {
		log.Panicln(err)
	}
	querydb.CreateAnxTable(ctx, conn)

	// Создаём объект бота, передаём токен.
	// Библиотека сама настроит URL, проверит токен и т.п.
	err = godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используем переменные окружения системы")
	}

	// 2. Получаем токен из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен в .env или переменных окружения")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		// Если не удалось создать бота (неправильный токен, нет сети и т.д.) — падаем.
		panic(err)
	}

	// Включаем debug-режим: библиотека будет писать в лог подробную информацию
	// о запросах/ответах к Telegram (удобно при отладке).
	bot.Debug = true

	// Логируем, что бот успешно авторизовался, и печатаем его username.
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создаём "настройку" получения обновлений.
	// Параметр 0 означает, что мы хотим получать все новые апдейты, начиная с текущего.
	u := tgbotapi.NewUpdate(0)

	// Timeout в секундах: сервер Telegram может "долго опрашиваться" (long polling),
	// чтобы не крутить цикл слишком часто. 60 секунд — стандартное значение.
	u.Timeout = 60

	// Получаем канал обновлений.
	// Внутри библиотека запускает цикл getUpdates и сама управляет offset.
	updates := bot.GetUpdatesChan(u)

	// Необязательный шаг: даём немного времени (500 мс),
	// чтобы в канал успели прилететь старые накопившиеся апдейты.
	time.Sleep(time.Millisecond * 500)

	// Очищаем канал обновлений, если не хотим обрабатывать "старый хвост" сообщений,
	// которые накопились до запуска бота.
	updates.Clear()

	// Основной цикл обработки входящих обновлений.
	// range по каналу updates будет получать каждый новый update, пока бот работает.

	for update := range updates {
		// Некоторые апдейты могут не содержать Message (например, callback_query, edited_message и т.п.).
		if update.Message != nil {

			router.HandleMessage(bot, &update, ctx, conn)

		} else if update.CallbackQuery != nil {
			router.HandleCallback(bot, &update, ctx, conn)
		}

	}
}
