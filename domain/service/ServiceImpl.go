package service

import (
	"log"
	"strings"
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/domain/entity"
	"ws_notifications_email/infrastructure/persistence"
	"ws_notifications_email/infrastructure/repository"
	"ws_notifications_email/utils"

	gmailgo "github.com/andresgaviria24/gmailgo"
)

const (
	ERROR      = "ERROR"
	SUCCESSFUL = "SUCCESSFUL"
)

type ServiceImpl struct {
	emailRepository repository.EmailRepository
	emailsConfs     *gmailgo.EmailService
}

func InitServiceImpl() *ServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ServiceImpl{
		emailRepository: dbHelper.EmailRepository,
		emailsConfs:     dbHelper.EmailsConfs,
	}
}

func (cd *ServiceImpl) SendEmail(email dto.Email) *dto.Response {

	cd.emailsConfs.Email.Body = email.Body
	cd.emailsConfs.Email.AddRecipients(email.To...)
	cd.emailsConfs.Email.Subject = email.Subject

	err := cd.emailsConfs.Email.Send(cd.emailsConfs.Auth)

	log := entity.Log{
		To:      strings.Join(email.To, ";"),
		System:  email.System,
		Body:    email.Body,
		Date:    utils.TimeNow(),
		Status:  SUCCESSFUL,
		Subject: email.Subject,
	}

	if err != nil {

		log.Status = ERROR
		log.Error = err.Error()

		cd.emailRepository.Log(log)

		return &dto.Response{
			Status:      400,
			Description: "error to send email",
		}
	}

	cd.emailRepository.Log(log)

	return &dto.Response{
		Status:      200,
		Description: "successful sent",
	}

}
