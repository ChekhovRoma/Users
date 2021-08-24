package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"users/models"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=romax password=mypassword dbname=romax sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable. Wait for %d sec. \n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
