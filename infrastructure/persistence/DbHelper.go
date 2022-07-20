package persistence

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ws_notifications_email/infrastructure/repository"
)

type DbHelper struct {
	EmailRepository repository.EmailRepository
	emailLogDb      *gorm.DB
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
	}, nil
}
