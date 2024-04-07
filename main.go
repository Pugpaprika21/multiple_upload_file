package main

import (
	"log"

	"github.com/Pugpaprika21/db"
	uploadFileRouter "github.com/Pugpaprika21/pkg/upload_file/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	db, err := db.New().UseDB()
	if err != nil {
		panic(err)
	}

	uploadFileRouter.Register(e, db)

	e.Logger.Fatal(e.Start(":8006"))
}
