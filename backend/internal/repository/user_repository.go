package repository
import (
	"weel-backend/internal/domain"
	"gorm.io/gorm"
)
type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uint) error
	List(limit, offset int) ([]*domain.User, error)
	Count() (int64, error)
}
type userRepository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
func (r *userRepository) List(limit, offset int) ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}
func (r *userRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&domain.User{}).Count(&count).Error
	return count, err
}
