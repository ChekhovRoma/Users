package service

// тут должна быть описана модель юзера которую вернет метод гет?
// мы сейчас на уровне сервиса и репозиторий не должен знать ничего об уровне который
//выше него?
type UserRepository interface {
	Create(name, email, password, role string) (int, error)
	Get(email, password string) (string, error)
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
	return s.userRepo.Get(email, password)
}
