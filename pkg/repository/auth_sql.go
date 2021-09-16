package repository

import (
	"github.com/jinzhu/gorm"
	"users/models"
)

type AuthSql struct {
	db *gorm.DB
}

func NewAuthSql(db *gorm.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (r *AuthSql) CreateUser(user models.User) (int, error) {

	result := r.db.Create(&user)

	if err := result.Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}
