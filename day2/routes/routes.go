package routes

import (
	"day2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	v1.GET("/books", controllers.GetAllBooks)
	v1.GET("/books/:id", controllers.GetBook)
	v1.POST("/books", controllers.CreateBook)
	v1.PUT("/books/:id", controllers.UpdateBook)
	v1.DELETE("/books/:id", controllers.DeleteBook)

	v1.GET("/users", controllers.GetAllUsers)
	v1.GET("/users/:id", controllers.GetUser)
	v1.POST("/users", controllers.CreateUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
