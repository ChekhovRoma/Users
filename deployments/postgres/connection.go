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

	db, err := gorm.Open("postgres", "postgresql://romax:mypassword@postgresCont/romax?sslmode=disable")

	//db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	//	viper.GetString("host"),
	//	viper.GetString("username"),
	//	viper.GetString("password"),
	//	viper.GetString("dbname"),
	//	viper.GetString("port"),
	//	viper.GetString("sslmode")))

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
