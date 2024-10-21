package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int       `json:"id" gorm:"primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	Phone        string    `json:"phone" gorm:"unique"`
	Photo        string    `json:"photo"`
	LinkTreeURL  string    `json:"link_tree_url"`
	Bio          string    `json:"bio"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
}
