package service

import (
	"ws_notifications_email/domain/dto"
)

type Service interface {
	SendEmail(dto.Email) *dto.Response
}
