package main

import (
	"day7/database"
	"day7/internal/middleware"
	"day7/internal/server"

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
