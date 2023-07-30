package db

import (
	"fmt"

	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/config"
	"github.com/abhinandkakkadi/moviesgo-products-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Genre{})
	db.AutoMigrate(&domain.MovieStudio{})
	return db, dbErr

}
