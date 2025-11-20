package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort            string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	MinioEndpoint      string
	MinioAccessKey     string
	MinioSecretKey     string
	MinioUseSSL        bool
	MinioBucket        string
	MinioPublicBaseURL string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &Config{
			AppPort:            getEnv("APP_PORT", ":8080"),
			DBHost:             getEnv("DB_HOST", "127.0.0.1"),
			DBPort:             getEnv("DB_PORT", "3306"),
			DBUser:             getEnv("DB_USER", "root"),
			DBPassword:         getEnv("DB_PASSWORD", ""),
			DBName:             getEnv("DB_NAME", "nontonkuys"),
			MinioEndpoint:      getEnv("MINIO_ENDPOINT", "127.0.0.1:9000"),
			MinioAccessKey:     getEnv("MINIO_ACCESS_KEY", ""),
			MinioSecretKey:     getEnv("MINIO_SECRET_KEY", ""),
			MinioUseSSL:        getEnv("MINIO_USE_SSL", "false") == "true",
			MinioBucket:        getEnv("MINIO_BUCKET", "category-images"),
			MinioPublicBaseURL: getEnv("MINIO_PUBLIC_BASE_URL", ""),
		}
	})

	return cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		if defaultVal == "" {
			log.Printf("WARNING: env %s is empty", key)
		}
		return defaultVal
	}
	return val
}
