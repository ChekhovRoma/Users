package service

type UserRepository interface {
	Create(name, email, password, role string) (int, error)
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
