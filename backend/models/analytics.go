package model

import "time"

type Analytics struct {
	ID         int       `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	ClickCount int       `json:"click_count"`
	UserID     int       `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
