package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"users/models"
)

const (
	signatureKey = "Y17GJYH13Bhjbj22gvG2J4Vh"
	tokenTTL     = 12 * time.Hour
)

type UserRepository interface {
	Create(name, email, password, role string) (int, error)
	Get(email, password string) (models.User, error)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthorizationService struct {
	userRepo UserRepository
}

func NewAuthorizationService(userRepo UserRepository) *AuthorizationService {
	return &AuthorizationService{userRepo: userRepo}
}

func (s *AuthorizationService) Create(name, email, password, role string) (int, error) {
	return s.userRepo.Create(name, email, password, role)
}

func (s *AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := s.userRepo.Get(email, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signatureKey))
}
