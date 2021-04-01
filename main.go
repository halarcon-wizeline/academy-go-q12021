package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"

	"github.com/halarcon-wizeline/academy-go-q12021/service"
	"github.com/halarcon-wizeline/academy-go-q12021/usecase"
	"github.com/halarcon-wizeline/academy-go-q12021/controller"
	"github.com/halarcon-wizeline/academy-go-q12021/router"
)

func main() {

	// logger setup
	logger, err := createLogger()
	if err != nil || logger == nil {
		log.Fatal("creating logger: %w", err)
	}

	// Service
	service, _ := service.New(logger)

	// Usecase
	useCase := usecase.New(service)

	// Controllers
	controller := controller.New(useCase, render.New())

	// Setup application routes
	httpRouter := router.New(controller)

	var serverPort string = "8080"
	fmt.Println("Server listen at http://localhost" + ":" + serverPort)
	log.Fatal(http.ListenAndServe("localhost:"+serverPort, httpRouter))
}

func createLogger() (*logrus.Logger, error) {
	logLevel := "DEBUG"
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"log_level": logLevel,
		}).Error("parsing log_level")

		return nil, err
	}

	logger := logrus.New()
	logger.SetLevel(level)
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	return logger, nil
}
