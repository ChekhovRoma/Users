package main

import (
	"fmt"
	_ "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"os/signal"
	_ "users/configs"
)

func testMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello")
}

func main() {
	//srv := new(users.Server)
	//if err := srv.Run("8084"); err != nil {
	//	log.Fatalf("error occured while running http server: %s", err.Error())
	//}
	c := make(chan os.Signal, 0)
	signal.Notify(c)

	http.HandleFunc("/docker", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, World(From Docker)!")
	})

	http.ListenAndServe(":8080", nil)

	s := <-c
	fmt.Println("Got signal:", s) //Got signal: terminated
	//r := mux.NewRouter()
	//r.HandleFunc("/test", testMethod).Methods("GET")
	//
	//fmt.Println("hello")

}
