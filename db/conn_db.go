package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"minijo/models"
)

func Init() *gorm.DB {
	//dsn := "postgres://minijo123:miata123@localhost:5432/minijo"
	dsn := "postgres://postgres:KrdZchlIfW@localhost:5432/minijo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Keyboard{}, &models.Review{}, &models.Like{})

	return db
}
