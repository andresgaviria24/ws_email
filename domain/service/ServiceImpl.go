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

	//emailsTo := strings.Join(email.To, ";")

	/*response := rest.Post(os.Getenv("API_URL")+"?apikey="+os.Getenv("API_KEY")+"&subject="+
	email.Subject+"&from="+os.Getenv("MAIL_USER")+"&fromName="+os.Getenv("MAIL_NAME")+"&to="+
	emailsTo+"&bodyHtml='"+email.Body+"'", nil)*/

	/*from := mail.NewEmail("no-reply", "noreply.message2022@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "warhammerjj28@gmail.com;andres_felipe_gaviria28@hotmail.com")*/

	/*m := mail.NewV3Mail()

	personalization := mail.NewPersonalization()*/

	sendEmail := new(dto.EmailBrevo)

	for _, m := range email.To {

		sendEmail.To = append(sendEmail.To, dto.InfoEmail{
			Name:  m,
			Email: m,
		})
		/*personalization.AddTos(&mail.Email{
			Name:    "",
			Address: m,
		})*/
	}

	//personalization.Subject = email.Subject

	//m.AddPersonalizations(personalization)

	/*from := mail.NewEmail(os.Getenv("MAIL_NAME"), os.Getenv("MAIL_USER"))
	content := mail.NewContent("text/html", email.Body)*/
	if len(email.AttachBase64) > 0 {
		/*attachment := mail.NewAttachment()
		attachment.SetContent(email.AttachBase64)
		attachment.SetType("application/pdf")
		if len(email.NameAttach) > 0 {
			email.NameAttach = uuid.New().String()
		}
		attachment.SetFilename(email.NameAttach + ".pdf")
		attachment.SetDisposition("attachment")
		m.AddAttachment(attachment)*/
		sendEmail.Attachment = append(sendEmail.Attachment, dto.Attachment{
			Content: email.AttachBase64,
			Name:    email.NameAttach + ".pdf",
		})
	}
	/*m.SetFrom(from)
	m.AddContent(content)*/

	/*request := sendgrid.GetRequest(os.Getenv("API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)*/

	sendEmail.Subject = email.Subject
	/*sendEmail.HTMLContent = html.EscapeString(email.Body)
	htmlString := strings.ReplaceAll(email.Body, `"`, `\"`)
	/*htmlString = strings.ReplaceAll(htmlString, "\n", " ")
	htmlString = strings.ReplaceAll(htmlString, `\'`, `\'`)
	email.Body = htmlString
	sendEmail.HTMLContent = htmlString*/
	//htmlString := strings.ReplaceAll(email.Body, `"`, `\"`)
	sendEmail.HTMLContent = strings.ReplaceAll(email.Body, `"`, `\"`)
	//sendEmail.HTMLContent = email.Body

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
