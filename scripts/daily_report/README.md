# Как пользоваться скриптом

🚀 Как запустить по cron:
- Скомпилируйте программу: 
    go build -o dailyreport dailyreport.go
- Настройте cron для запуска каждый день в 10 утра с понедельника по пятницу (crontab -e):
    0 10 * * 1-5 /path/to/dailybot