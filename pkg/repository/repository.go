package repository

import (
	"github.com/jinzhu/gorm"
	"users/models"
)

type UserRepository interface {
	Create(name, email, password, role string) (int, error)
	GetByCredentials(email, password string) (models.User, error)
	Get(id int) (models.User, error)
	Update(user models.User) (models.User, error)
}

type Repositories struct {
	UserRepo UserRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUser(db),
	}
}
