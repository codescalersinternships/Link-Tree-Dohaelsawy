package repository

import (
	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
)

func (db *DbInstance) GetUserID(u *model.User, id int) (err error) {
	if err := db.DB.Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) GetUserEmail(u *model.User, email string) (err error) {
	if err := db.DB.Where("email = ?", email).First(u).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) GetUserUsername(u *model.User, username string) (err error) {
	if err := db.DB.Where("username = ?", username).First(u).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) AddNewUser(u *model.User) (err error) {
	if err = db.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) PutOneUser(u *model.User, id int) (err error) {
	if err = db.DB.Model(&model.User{}).Where("id = ?", id).Updates(model.User{FirstName: u.FirstName, LastName: u.LastName,Phone: u.Phone, Bio: u.Bio, Image: u.Image}).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) DeleteUser(u *model.User, id int) (err error) {
	db.DB.Where("id = ?", id).Delete(u)
	return nil
}
