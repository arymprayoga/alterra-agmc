package controllers

import (
	"net/http"
	"strconv"
	"time"

	"day4/models"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Books{}
	idx   = 1
)

func GetAllBooks(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func CreateBook(ctx echo.Context) error {
	b := &models.Books{
		ID: idx,
	}
	if err := ctx.Bind(b); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"book":   nil,
		})
	}
	if err := ctx.Validate(b); err != nil {
		return err
	}
	b.CreatedAt = time.Now()
	books[b.ID] = b
	idx++
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
		"books":  books,
	})
}

func GetBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"book":   nil,
		})
	}

	if _, isPresent := books[id]; !isPresent {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "No Data Found",
			"book":   nil,
		})
	}

	return ctx.JSON(http.StatusOK, books[id])
}

func UpdateBook(ctx echo.Context) error {
	b := new(models.Books)
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.Bind(b); err != nil {
		return err
	}
	if _, isPresent := books[id]; !isPresent {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "No Data Found",
			"book":   nil,
		})
	}
	if err := ctx.Validate(b); err != nil {
		return err
	}
	books[id].Title = b.Title
	books[id].Author = b.Author
	books[id].Price = b.Price
	books[id].UpdatedAt = time.Now()
	return ctx.JSON(http.StatusOK, books[id])
}

func DeleteBook(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if _, isPresent := books[id]; !isPresent {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"status": "No Data Found",
			"book":   nil,
		})
	}
	delete(books, id)
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "Deleted",
	})
}
