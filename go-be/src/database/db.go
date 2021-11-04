package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hello/src/models"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Connect () {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,          // Disable color
		},
	)
	
	// Globally mode
	DB, err = gorm.Open(mysql.Open("root:root@tcp(mysql_go)/go_tutorial?parseTime=true"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Could not connect with database")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Product{}, models.Link{}, models.Order{}, models.OrderItem{})
}