package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"auth_svc/pkg/models"
)

type DBConnection struct {
	Conn *gorm.DB
}

func InitializeDBConnection(dsn string) DBConnection {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return DBConnection{Conn: db}

}
