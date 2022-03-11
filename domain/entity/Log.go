package entity

type Log struct {
	To      string `gorm:"NULL;SIZE=50;COLUMN:to"`
	System  string `gorm:"NULL;SIZE=50;COLUMN:system"`
	Id      int64  `gorm:"NULL;COLUMN:id"`
	Body    string `gorm:"NULL;SIZE=50;COLUMN:body"`
	Subject string `gorm:"NULL;SIZE=50;COLUMN:subject"`
	Date    string `gorm:"NULL;SIZE=50;COLUMN:date"`
	Status  string `gorm:"NULL;SIZE=50;COLUMN:status"`
	Error   string `gorm:"NULL;SIZE=50;COLUMN:error"`
}

func (Log) TableName() string {
	return "log"
}
