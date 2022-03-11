package persistence

import (
	"ws_notifications_email/domain/entity"

	"gorm.io/gorm"
)

type EmailRepositoryImpl struct {
	db *gorm.DB
}

func (upr *EmailRepositoryImpl) Log(log entity.Log) error {

	if err := upr.db.Create(&log).Error; err != nil {
		return err
	}
	return nil
}
