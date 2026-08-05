package models

import "time"

type Task struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Status   string `gorm:"default:PENDING"`
	Progress uint   `gorm:"default:0"`

	RetryCount int `gorm:"default:0"` // NEW

	Result string
	Error  string
	
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}
