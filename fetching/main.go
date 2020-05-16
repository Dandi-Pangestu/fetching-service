package main

import (
	_middleware "efishery.com/micro/middleware"
	"efishery.com/micro/router"
	"efishery.com/micro/shared/config"
	"efishery.com/micro/shared/utils"
	"github.com/labstack/echo/v4"
)

func init() {
	config.InitConfig()
}

func main() {
	middleware := _middleware.HTTPMiddleware()

	e := echo.New()
	e.Use(middleware.CORS)

	jwtService := _middleware.NewJWTService()
	authMiddleware := _middleware.NewAuthMiddleware(jwtService)
	redisClient := config.InitRedis()
	restService := utils.NewRestyRestService()

	router.NewFetchingRouter(e, authMiddleware, redisClient, restService)

	defer func() {
		if redisClient != nil {
			_ = redisClient.Close()
		}
	}()

	config.Startup(e)
}
