package repository

import (
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(email, password string) (int, error)
}

type Repositories struct {
	UserRepo UserRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUser(db),
	}
}
