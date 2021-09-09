package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"time"
	"users/models"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
	//	viper.GetString("username"),
	//	viper.GetString("password"),
	//	viper.GetString("dbname"),
	//	viper.GetString("sslmode"))
	//
	//log.Println(str)
	//db, err := gorm.Open("postgres", "user=romax password=mypassword dbname=romax sslmode=disable")

	db, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("sslmode")))

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	//if testDb != nil {
	//	log.Fatal("error initializing configs: ", testDb)
	//}
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

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
