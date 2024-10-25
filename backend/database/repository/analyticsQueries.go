package repository

import model "github.com/codescalersinternships/Link-Tree-Dohaelsawy/backend/models"

func (db *DbInstance) GetAllAnalyticsForUser(a *[]model.Analytics, user_id int) (err error) {
	if err := db.DB.Where("user_id = ?", user_id).Find(a).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) AddNewVisitor(a *model.Analytics) (err error) {
	if err = db.DB.Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) UpdateAnalytics(a *model.Analytics, id int) (err error) {
	if err = db.DB.Save(a).Error; err != nil {
		return err
	}
	return nil
}

func (db *DbInstance) GetAnalyticsForGuestUsername(a *model.Analytics, guestUsername string, userId int) (err error) {
	if err := db.DB.Where(&model.Analytics{GuestUsername: guestUsername, UserID: userId}).First(a).Error; err != nil {
		return err
	}
	return nil
}
