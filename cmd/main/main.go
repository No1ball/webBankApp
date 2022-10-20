package main

import (
	"github.com/No1ball/webBankApp/internal/config"
	"github.com/No1ball/webBankApp/internal/handlers"
	"github.com/No1ball/webBankApp/internal/server"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Error on initConfig: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error on godotenv: %s", err.Error())
	}

	handlers := handlers.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurated while running http server: %s", err.Error())
	}
}
