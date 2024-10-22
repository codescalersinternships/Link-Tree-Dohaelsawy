package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	ID         int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	Url        string    `json:"url" gorm:"unique"`
	UserID     int      `json:"user_id"`
	ClickCount int       `json:"click_count"`
	CreatedAt  time.Time `json:"created_at"`
}
