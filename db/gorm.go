package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"finalproject/server/models"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "P@$$w0rd"
	dbname = "MyGram"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	return db
}
