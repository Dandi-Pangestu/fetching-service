package router

import (
	"efishery.com/micro/httphandler"
	"efishery.com/micro/middleware"
	"efishery.com/micro/shared/utils"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
)

func NewFetchingRouter(e *echo.Echo, authMiddleware *middleware.AuthMiddleware, redisClient *redis.Client,
	restService utils.RestService) {

	handler := httphandler.FetchingHandler{
		RedisClient: redisClient,
		RestService: restService,
	}

	g := e.Group("/api/fetching")
	g.Use(authMiddleware.JWTAuth)

	g.GET("/fetch", handler.Fetch)
	g.GET("/claim-token", handler.ClaimToken)
	g.GET("/aggregate", handler.Aggregate, middleware.IsAdmin)
}
