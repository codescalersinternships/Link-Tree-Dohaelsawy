package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Image        string `json:"image"`
	LinkTreeURL  string `json:"link_tree_url"`
	Bio          string `json:"bio"`
	Token        string `json:"token"`
	CreatedAt time.Time `gorm:"created_at"`
}
