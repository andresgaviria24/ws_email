package service

import (
	"log"
	"strings"
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/domain/entity"
	"ws_notifications_email/infrastructure/persistence"
	"ws_notifications_email/infrastructure/repository"
	"ws_notifications_email/utils"

	"github.com/go-resty/resty/v2"
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

	sendEmail := new(dto.EmailBrevo)

	for _, m := range email.To {

		sendEmail.To = append(sendEmail.To, dto.InfoEmail{
			Name:  m,
			Email: m,
		})
	}

	if len(email.AttachBase64) > 0 {
		sendEmail.Attachment = append(sendEmail.Attachment, dto.Attachment{
			Content: email.AttachBase64,
			Name:    email.NameAttach + ".pdf",
		})
	}
	sendEmail.Subject = email.Subject
	htmlString := strings.ReplaceAll(email.Body, "\n", " ")
	sendEmail.HTMLContent = htmlString

	sendEmail.Sender = dto.InfoEmail{
		Name:  "No-Reply",
		Email: "noreply.message2022@gmail.com",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("content-type", "application/json").
		SetHeader("api-key", os.Getenv("API_KEY_BREVO")).
		SetHeader("accept", "application/json").
		SetBody(sendEmail).
		Post("https://api.brevo.com/v3/smtp/email")

	log.Println("status code", resp.Status())

	log := entity.Log{
		To:      strings.Join(email.To, ";"),
		System:  email.System,
		Body:    sendEmail.HTMLContent,
		Date:    utils.TimeNow(),
		Status:  SUCCESSFUL,
		Subject: email.Subject,
	}

	/*if resp != nil {

		if resp.StatusCode() != 200 {
			log.Status = ERROR
			log.Error = err.Error()

			cd.emailRepository.Log(log)

			return &dto.Response{
				Status:      400,
				Description: "error to send email",
			}
		}
	}*/

	if err != nil {
		log.Status = ERROR
		log.Error = err.Error()

		cd.emailRepository.Log(log)

		return &dto.Response{
			Status:      400,
			Description: "error to send email",
		}
	} else {
		//fmt.Println(resp.StatusCode())
		//fmt.Println(sendEmail)
	}

	cd.emailRepository.Log(log)

	return &dto.Response{
		Status:      200,
		Description: "successful sent",
	}

}
