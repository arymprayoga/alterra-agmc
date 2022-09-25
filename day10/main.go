package main

import (
	"day10/database"
	"day10/internal/middleware"
	"day10/internal/server"

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
