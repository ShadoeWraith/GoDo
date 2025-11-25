package db

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AutoMigrate(models ...any) error {

	if DB == nil {
		return errors.New("database connection is not initialized")
	}

	fmt.Println("Starting database migrations...")

	err := DB.AutoMigrate(models...)

	if err != nil {
		fmt.Printf("Database migration FAILED: %v\n", err)
		return err
	}

	fmt.Println("Database migrations completed successfully.")
	return nil
}

func getDSN() string {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

func InitGORM(models ...any) error {
	dsn := getDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database using GORM: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto-migrate database schema: %w", err)
	}

	DB = db
	fmt.Println("Successfully connected to PostgreSQL using GORM!")
	return nil
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()

		if err == nil {
			sqlDB.Close()
		}
	}
}
