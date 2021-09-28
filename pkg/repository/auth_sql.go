package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type (
	UserRepo struct {
		db *gorm.DB
	}

	User struct {
		ID        int    `json:"id" gorm:"primary_key"`
		Name      string `json:"name" binding:"required" gorm:"type:varchar(255);NOT NULL"`
		Email     string `json:"email" binding:"required" gorm:"type:varchar(255);NOT NULL;UNIQUE"`
		Password  string `json:"password" binding:"required" gorm:"type:varchar(50);NOT NULL"`
		Role      string `json:"role" gorm:"type:varchar(100);NOT NULL"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Users []User
)

func NewUser(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(name, email, password, role string) (int, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}

	result := r.db.Create(&user)
	if err := result.Error; err != nil {
		return 0, fmt.Errorf("user repo: create: %w", err)
	}

	return user.ID, nil
}
