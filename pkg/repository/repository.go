package repository

import (
	"github.com/jinzhu/gorm"
	"users/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
	}
}
