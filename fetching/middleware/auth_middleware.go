package middleware

import (
	"efishery.com/micro/shared/domains"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Authable Authable
}

func NewAuthMiddleware(authable Authable) *AuthMiddleware {
	return &AuthMiddleware{
		Authable: authable,
	}
}

func (m *AuthMiddleware) JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		authHeader := context.Request().Header.Get("Authorization")
		if authHeader == "" {
			err := domains.ErrUnauthenticate
			return context.JSON(http.StatusUnauthorized, domains.ResponseError{Message: err.Error()})
		}

		arrayAuthHeader := strings.Split(authHeader, " ")
		if len(arrayAuthHeader) != 2 {
			err := domains.ErrUnauthenticate
			return context.JSON(http.StatusUnauthorized, domains.ResponseError{Message: err.Error()})
		}

		if strings.ToLower(arrayAuthHeader[0]) != "bearer" {
			err := domains.ErrUnauthenticate
			return context.JSON(http.StatusUnauthorized, domains.ResponseError{Message: err.Error()})
		}

		tokenString := arrayAuthHeader[1]
		customClaim, err := m.Authable.Decode(tokenString)

		if (customClaim != nil) && err == nil {
			context.Set("id", customClaim.ID)
			context.Set("name", customClaim.Name)
			context.Set("phone", customClaim.Phone)
			context.Set("role", customClaim.Role)

			return next(context)
		}

		err = domains.ErrUnauthenticate
		return context.JSON(http.StatusUnauthorized, domains.ResponseError{Message: err.Error()})
	}
}
