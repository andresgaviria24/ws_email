package persistence

import (
	"fmt"
	"os"
	"time"

	gmailgo "github.com/andresgaviria24/gmailgo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ws_notifications_email/infrastructure/repository"
)

type DbHelper struct {
	EmailRepository repository.EmailRepository
	emailLogDb      *gorm.DB
	EmailsConfs     *gmailgo.EmailService
}

func InitDbHelper() (*DbHelper, error) {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=America/Bogota", host, user, pass, name, port)

	emailLogDb, err := gorm.Open(postgres.Open(config), &gorm.Config{Logger: logger.Default.LogMode(logger.Info), SkipDefaultTransaction: true})

	if err != nil {
		panic(err)
	}

	sqlDB, err := emailLogDb.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(600)
	sqlDB.SetMaxOpenConns(0)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	emailLogDb.AutoMigrate()
	return &DbHelper{
		EmailRepository: &EmailRepositoryImpl{emailLogDb},
		emailLogDb:      emailLogDb,
		EmailsConfs:     AuthGmail(),
	}, nil
}

func AuthGmail() *gmailgo.EmailService {

	a := gmailgo.Email{
		From:        os.Getenv("MAIL_USER"),
		Password:    os.Getenv("MAIL_PASS"),
		Name:        os.Getenv("MAIL_NAME"),
		ContentType: "text/html; charset=utf-8",
	}

	auth, err := a.Auth()
	if err != nil {
		return nil
	}

	return &gmailgo.EmailService{
		Auth:  auth,
		Email: &a,
	}
}
