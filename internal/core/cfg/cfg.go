package cfg

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type Config struct {
	GrpcPort            string
	AppBundleSecret     string // Secret key (32 bytes) to decrypt tenants db data
	PostgresDatabaseDSN string // admin db creds
}

var (
	once     sync.Once
	instance *Config
)

func Inst() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env file found, loading from OS environment variables.")
		}

		instance = &Config{
			GrpcPort:            getEnv("GRPC_PORT", "50051"),
			AppBundleSecret:     getEnv("APP_BUNDLE_SECRET", "12345678901234567890123456789012"),
			PostgresDatabaseDSN: getEnv("POSTGRES_DB_DSN", "postgres://pieceouser:pieceopassword@localhost:5432/atrace.tracker?sslmode=disable"),
		}
	})
	return instance
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
