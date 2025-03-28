package db

import (
	"fmt"

	config "github.com/joejosephvarghese/message/server/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		fmt.Println("❌ Database connection failed:", err)
		return nil, err
	}

	// Check if DB connection is working
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("❌ Failed to get database instance:", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("❌ Database ping failed:", err)
		return nil, err
	}

	fmt.Println("✅ Database connected successfully!")

	return db, nil
}
