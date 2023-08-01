package db

import (
	"fmt"

	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-user-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// dbName := "users"
	// err := db.Exec("CREATE DATABASE " + dbName).Error
	// _ = err

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Address{})
	return db, dbErr

}
