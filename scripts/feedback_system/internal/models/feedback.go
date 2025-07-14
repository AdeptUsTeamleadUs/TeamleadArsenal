package models

import "time"

// Feedback хранит ответ пользователя на опрос
type Feedback struct {
    ID         int       `db:"id"`
    UserID     int64     `db:"user_id"`
    Mood       int       `db:"mood"`          // Настроение (1-5)
    Workload   string    `db:"workload"`      // Загрузка (низкая/нормальная/высокая)
    Comment    string    `db:"comment"`       // Комментарий пользователя
    CreatedAt  time.Time `db:"created_at"`    // Время ответа
}