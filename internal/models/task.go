package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Status      string `gorm:"default:PENDING"`
	Result      string
	Error       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}
