package middleware

import (
	"day6/pkg/util/validator"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Validator = validator.Validator
}
