package models

import "time"

type User struct {
	ID             int       `json:"id" gorm:"primary_key"`
	Name           string    `json:"name" binding:"required" gorm:"type:varchar(255);NOT NULL"`
	Email          string    `json:"email" binding:"required" gorm:"type:varchar(255);NOT NULL;UNIQUE"`
	Password       string    `json:"password" binding:"required" gorm:"type:varchar(50);NOT NULL"`
	Role           string    `json:"role" gorm:"type:varchar(100);NOT NULL"`
	RefreshToken   string    `json:"refresh_token" gorm:"type:varchar(255)"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Users []User
