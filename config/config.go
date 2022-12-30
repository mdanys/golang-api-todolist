package config

import (
	"log"
	"os"
	"strconv"
	// "github.com/joho/godotenv"
)

type AppConfig struct {
	DBPort     int
	DBUser     string
	DBPass     string
	DBHost     string
	DBName     string
}

func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("cannot run configuration setup")
		return nil
	}

	return cfg
}

func initConfig() *AppConfig {
	var app AppConfig

	// godotenv.Load("config.env")

	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatal("error parse db port")
		return nil
	}
	app.DBPort = port

	app.DBUser = os.Getenv("MYSQL_USER")
	app.DBPass = os.Getenv("MYSQL_PASSWORD")
	app.DBHost = os.Getenv("MYSQL_HOST")
	app.DBName = os.Getenv("MYSQL_DBNAME")

	return &app
}
