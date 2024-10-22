package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	ID         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Url        string    `json:"url" gorm:"unique"`
	UserID     uint      `json:"user_id"`
	ClickCount int       `json:"click_count"`
	CreatedAt  time.Time `json:"created_at"`
}
