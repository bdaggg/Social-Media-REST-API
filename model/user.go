package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Username string    `json:"username" gorm:"uniqueIndex"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// Users struct
type Users struct {
	User []User `json:"user"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
