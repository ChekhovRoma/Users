package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"users/models"
)

func main() {
	//c := make(chan os.Signal, 0)
	//signal.Notify(c)
	//
	//fmt.Println("connecting")
	//// these details match the docker-compose.yml file.
	//postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	"postgres", 5432, "romax", "mypassword", "romax")
	//db, err := sql.Open("postgres", postgresInfo)
	//if err != nil {
	//	panic(err)
	//}
	////defer db.Close()
	//
	//start := time.Now()
	//for db.Ping() != nil {
	//	if start.After(start.Add(3 * time.Second)) {
	//		fmt.Println()
	//		panic("failed to connect after 3 secs.")
	//	}
	//}
	//fmt.Println("connected:", db.Ping() == nil)
	//_, err = db.Exec(`DROP TABLE IF EXISTS COMPANY;`)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec(`CREATE TABLE COMPANY (ID INT PRIMARY KEY NOT NULL, NAME text);`)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("table company is created")
	//fmt.Println("app started")
	//s := <-c
	//fmt.Println("Got signal:", s) //Got signal: terminated

	// @Shilin так не работает :(
	dbTest, err := gorm.Open("postgres", "user=romax password=mypassword dbname=romax sslmode=disabled")
	if err != nil {
		fmt.Println()
		panic(err)
	}
	dbTest.AutoMigrate(&models.User{})

	// @Shilin  так тоже не работает :(
	//postgres.Init()
}
