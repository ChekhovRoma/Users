package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//for docker
	//db, err := gorm.Open("postgres", "postgres://romax:mypassword@postgres-db/romax?sslmode=disable")

	//for local test
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("host"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("port"),
		viper.GetString("sslmode")))

	if err != nil {
		logrus.Println("connection error: ", err)
	}

	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		logrus.Println("db isn't init. start init...")
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			logrus.Printf("database is unavailable. wait for %d sec. \n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	logrus.Println("database is ready")
	return dbase
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
