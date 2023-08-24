package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	var err error
	dsn := "douyin:zjqxy@tcp(43.138.25.60:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//想要正确的处理time.Time,需要带上 parseTime 参数，
	//要支持完整的UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("err:", err.Error())
	}
}
