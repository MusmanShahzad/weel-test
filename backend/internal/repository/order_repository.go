package repository
import (
	"weel-backend/internal/domain"
	"gorm.io/gorm"
)
type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uint) (*domain.Order, error)
	GetByUserID(userID uint, limit, offset int) ([]*domain.Order, error)
	GetByUserIDWithFilters(userID uint, filters OrderFilters) ([]*domain.Order, error)
	Update(order *domain.Order) error
	Delete(id uint) error
}
type OrderFilters struct {
	Status             *string
	DeliveryPreference *string
	SortBy             string
	SortOrder          string
	Limit              int
	Offset             int
}
type orderRepository struct {
	db *gorm.DB
}
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}
func (r *orderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("User").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
func (r *orderRepository) GetByUserID(userID uint, limit, offset int) ([]*domain.Order, error) {
	var orders []*domain.Order
	err := r.db.Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}
func (r *orderRepository) GetByUserIDWithFilters(userID uint, filters OrderFilters) ([]*domain.Order, error) {
	var orders []*domain.Order
	query := r.db.Where("user_id = ?", userID)
	if filters.Status != nil && *filters.Status != "" {
		query = query.Where("status = ?", *filters.Status)
	}
	if filters.DeliveryPreference != nil && *filters.DeliveryPreference != "" {
		query = query.Where("delivery_preference = ?", *filters.DeliveryPreference)
	}
	sortBy := filters.SortBy
	if sortBy == "" {
		sortBy = "created_at"
	}
	sortOrder := filters.SortOrder
	if sortOrder == "" {
		sortOrder = "DESC"
	} else if sortOrder == "asc" {
		sortOrder = "ASC"
	} else {
		sortOrder = "DESC"
	}
	query = query.Order(sortBy + " " + sortOrder)
	if filters.Limit > 0 {
		query = query.Limit(filters.Limit)
	}
	if filters.Offset > 0 {
		query = query.Offset(filters.Offset)
	}
	err := query.Find(&orders).Error
	return orders, err
}
func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}
func (r *orderRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
