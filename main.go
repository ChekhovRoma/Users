package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"users/deployments/postgres"
	"users/models"
	"users/pkg/handler"
	"users/pkg/repository"
	"users/pkg/service"
)

func main() {
	var stopChan = make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	db := postgres.GetDB()
	db.AutoMigrate(&models.User{})

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
	log.Println("wait for sigterm/sigterm")
	<-stopChan // wait for SIGINT
}
