package db

import (
	"database/sql"
	"time"

	"github.com/AdeptUsTeamleadUs/feedback_system/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
    DB *sql.DB
}

// NewStorage создает новое подключение к базе данных
func NewStorage(dbPath string) (*Storage, error) {
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, err
    }
    // Создаем таблицу, если ее еще нет
    createTable := `
    CREATE TABLE IF NOT EXISTS feedback (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        mood INTEGER,
        workload TEXT,
        comment TEXT,
        created_at TIMESTAMP
    );`
    if _, err := db.Exec(createTable); err != nil {
        return nil, err
    }
    return &Storage{DB: db}, nil
}

// SaveFeedback сохраняет фидбек пользователя
func (s *Storage) SaveFeedback(f models.Feedback) error {
    query := `
    INSERT INTO feedback (user_id, mood, workload, comment, created_at)
    VALUES (?, ?, ?, ?, ?);`
    _, err := s.DB.Exec(query, f.UserID, f.Mood, f.Workload, f.Comment, f.CreatedAt.Format(time.RFC3339))
    return err
}

// GetAllFeedback возвращает все фидбеки (например, для дашборда)
func (s *Storage) GetAllFeedback() ([]models.Feedback, error) {
    rows, err := s.DB.Query("SELECT id, user_id, mood, workload, comment, created_at FROM feedback")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []models.Feedback
    for rows.Next() {
        var f models.Feedback
        var createdAt string
        if err := rows.Scan(&f.ID, &f.UserID, &f.Mood, &f.Workload, &f.Comment, &createdAt); err != nil {
            return nil, err
        }
        f.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
        result = append(result, f)
    }
    return result, nil
}