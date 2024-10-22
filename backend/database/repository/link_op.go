package repository

import (
	"fmt"

	model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"
)

func (db *DbInstance) GetAllLinks(l *[]model.Link) (err error) {
	if err = db.DB.Find(l).Error; err != nil {
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

func (db *DbInstance) PutOneLink(l *model.Link, id string) (err error) {
	fmt.Println(l)
	db.DB.Save(l)
	return nil
}

func (db *DbInstance) DeleteLink(l *model.Link, id string) (err error) {
	db.DB.Where("id = ?", id).Delete(l)
	return nil
}
