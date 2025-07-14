package main

import (
	"log"
	"time"

	tb "gopkg.in/telebot.v3"
)

const token = "YOUR_TELEGRAM_BOT_TOKEN"

func main() {
    pref := tb.Settings{
        Token:  token,
        Poller: &tb.LongPoller{Timeout: 10 * time.Second},
    }

    bot, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }

    bot.Handle("/start", func(c tb.Context) error {
        return c.Send("Привет! Каждую пятницу я буду задавать пару вопросов, чтобы понять твоё настроение в команде.")
    })

    questions := []string{
        "🟢 Оцени своё настроение в работе (1–5):",
        "🟡 Как ты оцениваешь свою загрузку? (низкая/нормальная/высокая):",
        "🔵 Что хотелось бы изменить в работе прямо сейчас?",
    }

    bot.Handle("/poll", func(c tb.Context) error {
        for _, q := range questions {
            if err := c.Send(q); err != nil {
                log.Println("Ошибка отправки вопроса:", err)
            }
        }
        return nil
    })

    log.Println("Бот запущен!")
    bot.Start()
}