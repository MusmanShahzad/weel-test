package service
import (
	"time"
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)
type AuthService interface {
	Login(email, password string) (*LoginResponse, error)
	GetCurrentUser(userID uint) (*domain.User, error)
}
type LoginResponse struct {
	Token string       `json:"token"`
	User  *domain.User `json:"user"`
}
type authService struct {
	userRepo   repository.UserRepository
	jwtService *JWTService
}
func NewAuthService(userRepo repository.UserRepository, jwtService *JWTService) AuthService {
	return &authService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}
func (s *authService) Login(email, password string) (*LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}
	now := time.Now()
	user.LastLogin = &now
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	token, err := s.jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
func (s *authService) GetCurrentUser(userID uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
