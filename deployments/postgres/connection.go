package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	//db, err := gorm.Open("postgres", "postgresql://romax:mypassword@postgres/romax?sslmode=disable")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("host"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("port"),
		viper.GetString("sslmode")))

	if err != nil {
		fmt.Print("\n connection error: ", err)
		//log.Fatal(err)
	}

	//db.AutoMigrate(&models.User{})
	//if testDb != nil {
	//	log.Fatal("error initializing configs: ", testDb)
	//}
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		fmt.Print("DB isn't init. start init... \n ")
		dbase = Init()
		fmt.Print("\n error while init: ", dbase.Error)
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable. Wait for %d sec. \n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	fmt.Print("Database is ready \n")
	return dbase
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
