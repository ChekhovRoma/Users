package models

import "time"

type User struct {
	ID        string `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);NOT NULL"`
	Email     string `gorm:"type:varchar(255);NOT NULL;UNIQUE"`
	Password  string `gorm:"type:varchar(50);NOT NULL"`
	Role      string `gorm:"type:varchar(100);NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User
