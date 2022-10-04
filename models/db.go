package models

import (
	"time"

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
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Custom     bool   `gorm:"not null"`
	Photo_path string `gorm:"default:/static/img/placeholder.png"`
	PcbID      int    `gorm:"default:null"`
	CaseID     int    `gorm:"default:null"`
	SwitchID   int    `gorm:"default:null"`
	KeycapsID  int    `gorm:"default:null"`
}

type Users struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
}

type Reviews struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Keyboard   Keyboard
	KeyboardID int
	User       Users
	Date       time.Time
	UserID     int
	Stars      int `gorm:"check:stars >= 0,stars <= 5"`
	Likes      int `gorm:"check:likes >= 0"`
}
