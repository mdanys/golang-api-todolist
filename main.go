package main

import (
	"golang/config"
	ad "golang/features/activity/delivery"
	ar "golang/features/activity/repository"
	as "golang/features/activity/service"
	td "golang/features/todo/delivery"
	tr "golang/features/todo/repository"
	ts "golang/features/todo/service"
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

	tRepo := tr.New(db)
	tService := ts.New(tRepo)
	td.New(e, tService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":3030"))
}
