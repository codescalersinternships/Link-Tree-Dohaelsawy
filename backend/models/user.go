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
	Phone        string    `json:"phone"`
	Photo        string    `json:"photo"`
	LinkTreeURL  string    `json:"link_tree_url"`
	Bio          string    `json:"bio"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
}




// func GetAllBook(b *[]User) (err error) {
// 	if err = Config.DB.Find(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func AddNewBook(b *Book) (err error) {
// 	if err = Config.DB.Create(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetOneBook(b *Book, id string) (err error) {
// 	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func PutOneBook(b *Book, id string) (err error) {
// 	fmt.Println(b)
// 	Config.DB.Save(b)
// 	return nil
// }

// func DeleteBook(b *Book, id string) (err error) {
// 	Config.DB.Where("id = ?", id).Delete(b)
// 	return nil
// }