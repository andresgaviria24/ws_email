package persistence

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
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

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	emailLogDb, err := gorm.Open(mysql.Open(config), &gorm.Config{Logger: logger.Default.LogMode(logger.Info), SkipDefaultTransaction: true})

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
