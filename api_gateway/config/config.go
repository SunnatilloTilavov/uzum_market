package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const ()

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	RedisHost     string
	RedisPort     int
	RedisPassword string

	CatalogServiceHost string
	CatalogServicePort string
		
	OrderServiceHost string
	OrderServicePort string

	UserServiceHost string
	UserServicePort string
	
	AuthServiceHost string
	AuthServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "127.0.0.1"))
	c.RedisPort = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "3EEdwhDOfx"))

	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "localhost"))
	c.CatalogServicePort = cast.ToString(getOrReturnDefault("CATALOG_GRPC_PORT", "8082"))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("Order_SERVICE_HOST", "localhost"))
	c.OrderServicePort = cast.ToString(getOrReturnDefault("Order_GRPC_PORT", "8083"))
	
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_GRPC_PORT", "8084"))

	c.AuthServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.AuthServicePort = cast.ToString(getOrReturnDefault("USER_GRPC_PORT", "8085"))
	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
