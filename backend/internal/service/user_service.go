package service
import (
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
)
type UserService interface {
	CreateUser(req *CreateUserRequest) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUser(id uint, req *UpdateUserRequest) (*domain.User, error)
	DeleteUser(id uint) error
	ListUsers(limit, offset int) ([]*domain.User, int64, error)
}
type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}
type UpdateUserRequest struct {
	Email     *string `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
type userService struct {
	userRepo repository.UserRepository
}
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}
func (s *userService) CreateUser(req *CreateUserRequest) (*domain.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, ErrEmailExists
	}
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
func (s *userService) UpdateUser(id uint, req *UpdateUserRequest) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	if req.Email != nil && *req.Email != user.Email {
		existingUser, _ := s.userRepo.GetByEmail(*req.Email)
		if existingUser != nil {
			return nil, ErrEmailExists
		}
		user.Email = *req.Email
	}
	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userService) DeleteUser(id uint) error {
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return ErrUserNotFound
	}
	return s.userRepo.Delete(id)
}
func (s *userService) ListUsers(limit, offset int) ([]*domain.User, int64, error) {
	users, err := s.userRepo.List(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.userRepo.Count()
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
