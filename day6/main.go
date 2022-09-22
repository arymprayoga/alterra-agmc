package main

import (
	"day6/database"
	"day6/internal/middleware"
	"day6/internal/server"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()
	e := echo.New()
	middleware.Init(e)
	server.NewHttp(e)
	e.Logger.Fatal(e.Start(":8000"))
}
