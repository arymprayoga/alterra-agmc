package middleware

import (
	"day7/pkg/util/validator"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.Validator = validator.Validator
}
