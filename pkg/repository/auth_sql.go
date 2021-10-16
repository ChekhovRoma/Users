package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"users/models"
)

type (
	UserRepo struct {
		db *gorm.DB
	}

	User struct {
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

func (r *UserRepo) GetByCredentials(email, password string) (models.User, error) {
	var user models.User
	result := r.db.Where(&models.User{Email: email, Password: password}).Find(&user)

	return user, result.Error
}

func (r *UserRepo) Get(id int) (models.User, error) {
	var user models.User
	result := r.db.Where(&models.User{ID: id}).Find(&user)

	return user, result.Error
}

// Update todo кажется хуита
func (r *UserRepo) Update(user models.User) (models.User, error) {
	result := r.db.Save(user)

	return user, result.Error
}
