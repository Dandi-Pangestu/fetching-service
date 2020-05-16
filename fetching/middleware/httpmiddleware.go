package middleware

import (
	"github.com/labstack/echo/v4"
)

type HTTPCreateMiddleware struct{}

func HTTPMiddleware() *HTTPCreateMiddleware {
	return &HTTPCreateMiddleware{}
}

func (m *HTTPCreateMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
