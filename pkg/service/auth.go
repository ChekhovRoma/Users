package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"time"
	"users/models"
	"users/pkg/handler"
)

const (
	signatureKey = "Y17GJYH13Bhjbj22gvG2J4Vh"
	tokenTTL     = 12 * time.Hour
)

type UserRepository interface {
	Create(name, email, password, role string) (int, error)
	GetByCredentials(email, password string) (models.User, error)
	Get(id int) (models.User, error)
	Update(user models.User) (models.User, error)
}

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthorizationService struct {
	userRepo UserRepository
	tm       TokenManager
}

func NewAuthorizationService(userRepo UserRepository, tm TokenManager) *AuthorizationService {
	return &AuthorizationService{
		userRepo: userRepo,
		tm:       tm,
	}
}

func (s *AuthorizationService) SignUp(name, email, password, role string) (int, error) {
	return s.userRepo.Create(name, email, password, role)
}

func (s *AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := s.userRepo.GetByCredentials(email, password)
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

func (s *AuthorizationService) SignIn(ctx context.Context, email, password string) (handler.Tokens, error) {
	user, err := s.userRepo.GetByCredentials(email, password)
	if err != nil {
		return handler.Tokens{}, fmt.Errorf("get user: %w", err)
	}

	return s.CreateSession(ctx, user.ID)
}

func (s *AuthorizationService) CreateSession(ctx context.Context, userId int) (handler.Tokens, error) {
	var (
		res handler.Tokens
		err error
	)

	id := cast.ToString(userId)

	//todo wrap err
	res.AccessToken, err = s.tm.NewJWT(id, tokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tm.NewRefreshToken()
	if err != nil {
		return res, err
	}

	user, err := s.userRepo.Get(userId)
	if err != nil {
		return res, err
	}

	user.RefreshToken = res.RefreshToken
	user.TokenExpiredAt = time.Now().Add(tokenTTL)

	//todo is this naming huita?
	_, err = s.userRepo.Update(user)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s *AuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signatureKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
