package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/handlers"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/middlewares"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/routes"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/applicationData"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/database"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/mailer"
)

func main() {

	e := echo.New()

	err := godotenv.Load(".env")

	if err != nil {

		e.Logger.Fatal(err.Error())
	}

	db, err := database.ConnectMysql()

	if err != nil {
		e.Logger.Fatal(err.Error())
	}

	appMailer := mailer.NewMailer(e.Logger)

	h := handlers.Handler{
		Logger: e.Logger,
		Db:     db,
		Mailer: appMailer,
	}

	appMiddleware := middlewares.AppMiddleware{
		Db:     db,
		Logger: e.Logger,
	}

	appData := &applicationData.Application{
		Logger:        e.Logger,
		Server:        e,
		Handler:       h,
		AppMiddleware: appMiddleware,
	}

	routes.Route(appData)

	e.Use(middleware.Logger())

	domain := os.Getenv("APP_DOMAIN")
	port := os.Getenv("APP_PORT")

	appAddress := fmt.Sprintf("%s:%s", domain, port)

	e.Logger.Fatal(e.Start(appAddress))

}
