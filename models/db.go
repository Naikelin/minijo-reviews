package models

import (
	"gorm.io/gorm"
)

type Pcb struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}

type Case struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}

type Switch struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}

type Keycaps struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}

type Keyboard struct {
	gorm.Model
	ID       int     `gorm:"primaryKey"`
	Name     string  `gorm:"not null" json:"name"`
	UrlPhoto string  `gorm:"default: ''" json:"url_photo"`
	Stars    float64 `json:"stars"`
}

type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
}

type Review struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	KeyboardID  int    `gorm:"not null" json:"keyboard_id"`
	UserID      int    `gorm:"not null" json:"user_id"`
	Stars       int    `gorm:"check:stars >= 0; stars <= 5"`
	Description string `gorm:"default: ''" json:"description"`
	Likes       int    `gorm:"default: 0; check:likes >= 0"`
}

type Like struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	ReviewID int `json:"review_id"`
	UserID   int `json:"user_id"`
}
