## Установка и запуск

Перед сборкой и запуском убедитесь, что у вас установлен Go (1.19+).

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/yourusername/teamlead-motivation-dashboard.git
    cd teamlead-motivation-dashboard
    ```

2. Установите недостающие пакеты:

    ```bash
    go mod tidy
    ```

    Это подтянет все зависимости, включая:

    - [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) — драйвер SQLite для Go
    - [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) — web-фреймворк для интерфейса
    - [gopkg.in/telebot.v3](https://github.com/tucnak/telebot) — библиотека для Telegram-бота

3. Запустите Telegram-бота и веб-интерфейс в отдельных терминалах:

    ```bash
    go run cmd/bot/main.go
    # или
    go run cmd/web/main.go
    ```

4. Проверьте работу веб-интерфейса на [http://localhost:8080](http://localhost:8080)