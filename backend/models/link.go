package model

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	ID         int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	UserID     int      `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
