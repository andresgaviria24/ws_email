package repository

import (
	"ws_notifications_email/domain/entity"
)

type EmailRepository interface {
	Log(entity.Log) error
}
