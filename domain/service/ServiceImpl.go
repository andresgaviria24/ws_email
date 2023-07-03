package service

import (
	"fmt"
	"log"
	"os"
	"strings"
	"ws_notifications_email/domain/dto"
	"ws_notifications_email/domain/entity"
	"ws_notifications_email/infrastructure/persistence"
	"ws_notifications_email/infrastructure/repository"
	"ws_notifications_email/utils"

	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	//emailsTo := strings.Join(email.To, ";")

	/*response := rest.Post(os.Getenv("API_URL")+"?apikey="+os.Getenv("API_KEY")+"&subject="+
	email.Subject+"&from="+os.Getenv("MAIL_USER")+"&fromName="+os.Getenv("MAIL_NAME")+"&to="+
	emailsTo+"&bodyHtml='"+email.Body+"'", nil)*/

	/*from := mail.NewEmail("no-reply", "noreply.message2022@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "warhammerjj28@gmail.com;andres_felipe_gaviria28@hotmail.com")*/

	m := mail.NewV3Mail()

	personalization := mail.NewPersonalization()

	for _, m := range email.To {
		personalization.AddTos(&mail.Email{
			Name:    "",
			Address: m,
		})
	}

	personalization.Subject = email.Subject

	m.AddPersonalizations(personalization)

	from := mail.NewEmail(os.Getenv("MAIL_NAME"), os.Getenv("MAIL_USER"))
	content := mail.NewContent("text/html", email.Body)
	if len(email.AttachBase64) > 0 {
		attachment := mail.NewAttachment()
		attachment.SetContent(email.AttachBase64)
		attachment.SetType("application/pdf")
		if len(email.NameAttach) > 0 {
			email.NameAttach = uuid.New().String()
		}
		attachment.SetFilename(email.NameAttach + ".pdf")
		attachment.SetDisposition("attachment")
		m.AddAttachment(attachment)
	}
	m.SetFrom(from)
	m.AddContent(content)

	request := sendgrid.GetRequest(os.Getenv("API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)

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
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	cd.emailRepository.Log(log)

	return &dto.Response{
		Status:      200,
		Description: "successful sent",
	}

}
