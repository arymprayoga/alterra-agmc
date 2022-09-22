package auth

import (
	"net/http"
	"os"
	"time"

	"day7/internal/dto"
	"day7/internal/middleware"
	"day7/internal/repository"
	"day7/pkg/util/response"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Token(c echo.Context) error {
	payload := new(dto.AuthLoginRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, response.Failed{Success: false, Code: http.StatusBadRequest, Message: err.Error()})
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, response.Failed{Success: false, Code: http.StatusBadRequest, Message: err.Error()})
	}

	user, err := repository.FindByEmail(payload.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Failed{Success: false, Code: http.StatusUnauthorized, Message: "Incorrect email or password"})
	}

	claims := middleware.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "task-middleware",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(1) * time.Hour)),
		},
		ID:        user.ID,
		UserAgent: c.Request().Header.Get("user-agent"),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, false)
	}

	return c.JSON(http.StatusOK, &dto.AuthLoginResponse{Token: signedToken})
}
