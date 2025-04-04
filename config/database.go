package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gaelzamora/spent-one/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	fmt.Println("✅ Base de datos conectada con éxito")
	db.AutoMigrate(&domain.User{}, &domain.Spent{})
	return db
}
