package main

import (
	"net/http"

	"day3/config"
	"day3/middlewares"
	"day3/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	user, err = middlewares.CreateToken(1)
	config.InitDB()
	e := routes.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Logger.Fatal(e.Start(":8000"))
}
