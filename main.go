package main

import (
	"fmt"
	"log"
	"net/http"
)

func dockerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, World(From Docker)!")
}

func main() {
	//srv := new(users.Server)
	//if err := srv.Run("8084"); err != nil {
	//	log.Fatalf("error occured while running http server: %s", err.Error())
	//}
	//c := make(chan os.Signal, 0)
	//signal.Notify(c)

	http.HandleFunc("/docker", dockerHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	//s := <-c
	//fmt.Println("Got signal:", s) //Got signal: terminated
	//r := mux.NewRouter()
	//r.HandleFunc("/test", testMethod).Methods("GET")
	//
	//fmt.Println("hello")

}
