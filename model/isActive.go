package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type IsActive struct {
	gorm.Model
	UserID   uuid.UUID
	User     User `gorm:"foreignKey:UserID"`
	IsActive bool
}

type LogIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogOut struct {
	Username string `json"username"`
}
