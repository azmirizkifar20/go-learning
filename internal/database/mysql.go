package database

import (
	"fmt"
	"log"
	"sync"
	"time"

	"go-learning/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(func() {
		cfg := config.LoadConfig()

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)

		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		// set connection pool settings
		sqlDB, err := conn.DB()
		if err != nil {
			log.Fatalf("failed to get generic DB from gorm: %v", err)
		}

		sqlDB.SetMaxOpenConns(25)                 // maksimal koneksi terbuka ke DB
		sqlDB.SetMaxIdleConns(10)                 // maksimal koneksi idle di pool
		sqlDB.SetConnMaxLifetime(5 * time.Minute) // umur maksimal 1 koneksi

		db = conn
	})

	return db
}
