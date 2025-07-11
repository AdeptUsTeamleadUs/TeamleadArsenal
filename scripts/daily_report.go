package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	botToken = "YOUR_BOT_TOKEN"
	chatID   = "@your_team_chat"
)

func sendDailyReminder() error {
	message := "📝 Команда, доброе утро!\n\nПора написать статус на сегодня:\n- Что сделал вчера?\n- Что планирую сегодня?\n- Есть ли проблемы или блокеры?"

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	_, err := http.PostForm(apiURL, url.Values{
		"chat_id": {chatID},
		"text":    {message},
	})
	return err
}

func main() {
	if err := sendDailyReminder(); err != nil {
		fmt.Println("Ошибка отправки сообщения:", err)
	} else {
		fmt.Println("Напоминание успешно отправлено!")
	}
}