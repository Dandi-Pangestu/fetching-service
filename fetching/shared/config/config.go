package config

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.AddConfigPath("shared/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func Startup(e *echo.Echo) {
	host := viper.GetString("application.server.host")
	port := viper.GetInt("application.server.port")
	err := e.Start(fmt.Sprintf("%s:%d", host, port))
	log.Fatal(err)
}

func InitRedis() *redis.Client {
	instance := viper.GetString(fmt.Sprintf("application.resources.redis.instance"))
	port := viper.GetInt(fmt.Sprintf("application.resources.redis.port"))
	dbnumber := viper.GetInt(fmt.Sprintf("application.resources.redis.dbnumber"))
	password := viper.GetString(fmt.Sprintf("application.resources.redis.password"))

	if len(instance) == 0 {
		log.Fatalf("Redis instance is required")
	}

	if port < 0 {
		log.Fatalf("Redis port is required")
	}

	if dbnumber < 0 {
		log.Fatalf("Redis dbnumber is required")
	}

	addr := fmt.Sprintf("%s:%d", instance, port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbnumber,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Error while connect to redis server: %v", err)
	}

	return client
}
