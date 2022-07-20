package service

import (
	"log"
	"os"
	"strings"
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/domain/entity"
	"ws_notifications_email/infrastructure/persistence"
	"ws_notifications_email/infrastructure/repository"
	"ws_notifications_email/utils"

	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	ERROR      = "ERROR"
	SUCCESSFUL = "SUCCESSFUL"
)

type ServiceImpl struct {
	emailRepository repository.EmailRepository
	//emailsConfs     *gmailgo.EmailService
}

func InitServiceImpl() *ServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ServiceImpl{
		emailRepository: dbHelper.EmailRepository,
	}
}

func (cd *ServiceImpl) SendEmail(email dto.Email) *dto.Response {

	emailsTo := strings.Join(email.To, ";")

	response := rest.Post(os.Getenv("API_URL")+"?apikey="+os.Getenv("API_KEY")+"&subject="+
		email.Subject+"&from="+os.Getenv("MAIL_USER")+"&fromName="+os.Getenv("MAIL_NAME")+"&to="+
		emailsTo+"&bodyHtml='"+email.Body+"'", nil)

	log := entity.Log{
		To:      strings.Join(email.To, ";"),
		System:  email.System,
		Body:    email.Body,
		Date:    utils.TimeNow(),
		Status:  SUCCESSFUL,
		Subject: email.Subject,
	}

	if response.Err != nil {

		log.Status = ERROR
		log.Error = response.Err.Error()

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
