package model

import (
	//"mime/multipart"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	//File        multipart.File `json:"file,omitempty" validate:"required"` 
	ID1         uint `gorm:"primaryKey;autoIncrement"`
	Description string
	Username    string
	UserID      uuid.UUID
	User        User `gorm:"foreignKey:UserID"`
}

type PostRequest struct {
	Description string
	//Image       []byte
	Username string
}
