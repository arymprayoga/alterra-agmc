package server

import (
	"day10/internal/app/auth"
	"day10/internal/app/user"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo) {
	v1 := e.Group("/v1")
	user.Route(v1.Group("/users"))
	auth.Route(v1.Group("/auth"))
}
