package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"max/internal/models"
	"max/internal/config"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Account{}, 
		&models.Credit{}, 
		&models.PaymentSchedule{}, 
		&models.Transaction{}, 
		&models.User{}, 
		&models.Card{},
	)
}
