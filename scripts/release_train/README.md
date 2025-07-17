#### Telegram-уведомления

`notify_telegram.go` каждый день присылает в канал/группу обратный отсчёт до Code Freeze и Production Release.

1. Создайте Telegram-бота в @BotFather и скопируйте токен.
2. Добавьте бота в нужный канал / группу и выдайте ему право постить сообщения.
3. Узнайте `chat_id`: можно использовать @userinfobot или добавить бота в группу и посмотреть `chat_id` в ссылке «…start=<id>».
4. В файле `notify_telegram.go` замените `botToken`, `chatID`, `codeFreeze`, `releaseDay`.
5. Запускайте вручную или добавьте cron-задачу:

```bash
# daily at 09:00
0 9 * * * /usr/local/bin/go run /path/to/scripts/release_train/notify_telegram.go