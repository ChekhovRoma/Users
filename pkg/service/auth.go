package service

type UserRepository interface {
	Create(email, password string) (int, error)
}

type AuthorizationService struct {
	userRepo UserRepository
}

func NewAuthorizationService(userRepo UserRepository) *AuthorizationService {
	return &AuthorizationService{userRepo: userRepo}
}

func (s *AuthorizationService) Create(email, password string) (int, error) {
	return s.userRepo.Create(email, password)
}
