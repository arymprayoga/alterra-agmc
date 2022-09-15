package controllers

import (
	"net/http"
	"strconv"

	"github.com/arymprayoga/alterra-agmc/day2/lib/database"
	"github.com/arymprayoga/alterra-agmc/day2/models"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(ctx echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	users, e := database.GetSingleUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUser(ctx echo.Context) error {
	u := &models.Users{}

	if err := ctx.Bind(u); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"book":   nil,
		})
	}
	if err := ctx.Validate(u); err != nil {
		return err
	}

	users, e := database.CreateUser(*u)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func UpdateUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	u := models.Users{}
	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"book":   nil,
		})
	}

	users, err := database.UpdateUser(u, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func DeleteUser(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	e := database.DeleteUser(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
