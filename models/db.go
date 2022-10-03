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
	ID         int `gorm:"primaryKey"`
	Name       string
	Custom     bool
	Photo_path string
	Pcb        Pcb
	PcbID      int
	Case       Case
	CaseID     int
	Switch     Switch
	SwitchID   int
	Keycaps    Keycaps
	KeycapsID  int
}

type Users struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Name     string
}

type Reviews struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	Keyboard   Keyboard
	KeyboardID int
	User       Users
	Date       time.Time
	UserID     int
	Stars      int
	Likes      int
}
