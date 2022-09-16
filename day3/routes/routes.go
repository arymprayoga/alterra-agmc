package routes

import (
	"day3/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	jwtSecret := os.Getenv("JWT_SECRET")
	v1 := e.Group("/v1")
	v1.POST("/login", controllers.LoginUsers)

	v1Users := e.Group("/v1")
	v1Users.GET("/books", controllers.GetAllBooks)
	v1Users.GET("/books/:id", controllers.GetBook)
	v1Users.Use(middleware.JWT([]byte(jwtSecret)))
	v1Users.POST("/books", controllers.CreateBook)
	v1Users.PUT("/books/:id", controllers.UpdateBook)
	v1Users.DELETE("/books/:id", controllers.DeleteBook)

	v1Books := e.Group("/v1")
	v1Books.POST("/users", controllers.CreateUser)
	v1Books.Use(middleware.JWT([]byte(jwtSecret)))
	v1Books.GET("/users", controllers.GetAllUsers)
	v1Books.GET("/users/:id", controllers.GetUser)
	v1Books.PUT("/users/:id", controllers.UpdateUser)
	v1Books.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
