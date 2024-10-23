package repository

import (
	"fmt"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
)

func (db *DbInstance) GetAllLinksForUser(l *[]model.Link, user_id int) (err error) {
	if err := db.DB.Where("user_id = ?", user_id).Find(l).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) GetOneLink(l *model.Link, id int) (err error) {
	if err := db.DB.Where("id = ?", id).First(l).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) AddNewLink(l *model.Link) (err error) {
	if err = db.DB.Create(l).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) PutOneLink(l *model.Link, id int) (err error) {
	fmt.Println(l)
	if err = db.DB.Save(l).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) DeleteLink(l *model.Link, id int) (err error) {
	db.DB.Where("id = ?", id).Delete(l)
	return nil
}
