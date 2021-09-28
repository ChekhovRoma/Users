package repository

import (
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Create(name, email, password, role string) (int, error)
	Get(email, password string) (string, error)
}

type Repositories struct {
	UserRepo UserRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUser(db),
	}
}
