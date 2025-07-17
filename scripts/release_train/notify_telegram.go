package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	botToken  = "YOUR_TELEGRAM_BOT_TOKEN"   // токен вашего бота
	chatID    = "-1001234567890"            // ID канала/группы (со знаком минус) или @username
	codeFreeze = "2025-08-15"               // дата code-freeze
	releaseDay = "2025-08-30"               // дата прод-релиза
)

// daysLeft считает, сколько суток осталось до контрольной даты
func daysLeft(target string) int {
	t, _ := time.Parse("2006-01-02", target)
	return int(time.Until(t).Hours() / 24)
}

func main() {
	msg := fmt.Sprintf(
		"📢 *Release Train update*\n• %d days to *Code Freeze*\n• %d days to *Production Release*",
		daysLeft(codeFreeze),
		daysLeft(releaseDay),
	)

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	_, err := http.PostForm(apiURL, url.Values{
		"chat_id":    {chatID},
		"text":       {msg},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		fmt.Println("error sending message:", err)
		os.Exit(1)
	}
	fmt.Println("Telegram notification sent successfully")
}