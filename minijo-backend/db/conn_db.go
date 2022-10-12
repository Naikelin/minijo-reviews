package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"minijo/models"
)

func Init() *gorm.DB {
	user_db := os.Getenv("USER_DB")
	pass_db := os.Getenv("PASS_DB")
	host_db := os.Getenv("HOST_DB")
	port_db := os.Getenv("PORT_DB")
	database := os.Getenv("NAME_DB")
	dsn := "postgres://" + user_db + ":" + pass_db + "@" + host_db + ":" + port_db + "/" + database + "?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Keyboard{}, &models.Review{}, &models.Like{})

	return db
}
