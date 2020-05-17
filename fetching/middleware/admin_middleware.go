package middleware

import (
	"efishery.com/micro/shared/domains"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		role := context.Get("role").(string)
		if len(role) == 0 || role != "admin" {
			err := domains.ErrUnauthenticate
			return context.JSON(http.StatusUnauthorized, domains.ResponseError{Message: err.Error()})
		}

		return next(context)
	}
}
