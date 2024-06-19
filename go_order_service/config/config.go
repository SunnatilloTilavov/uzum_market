package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string
	Version     string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	ContentServiceHost string
	ContentGRPCPort    string

	PostgresMaxConnections int32
}

func Load() Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("not found .env file:", err)
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "catalog_service"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "sunnatillo"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1111"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "uzum_market"))
	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.ContentServiceHost = cast.ToString(getOrReturnDefaultValue("catalog_service_HOST", "localhost"))
	config.ContentGRPCPort = cast.ToString(getOrReturnDefaultValue("CONTENT_GRPC_PORT", ":8083"))
config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 1))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	return defaultValue
}