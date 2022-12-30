package main

import (
	"golang/config"
	ad "golang/features/activity/delivery"
	ar "golang/features/activity/repository"
	as "golang/features/activity/service"
	"golang/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	aRepo := ar.New(db)
	aService := as.New(aRepo)
	ad.New(e, aService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":3030"))
}
