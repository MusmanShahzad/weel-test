package service
import (
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
)
type OrderService interface {
	GetAISuggestions(req *GetAISuggestionsRequest) ([]domain.AISuggestedProduct, error)
	CreateOrder(userID uint, req *CreateOrderRequest) (*domain.Order, error)
	GetOrders(userID uint, filters *GetOrdersFilters) ([]*domain.Order, error)
	GetOrderByID(orderID, userID uint) (*domain.Order, error)
	UpdateOrder(orderID, userID uint, req *UpdateOrderRequest) (*domain.Order, error)
}
type GetAISuggestionsRequest struct {
	Summary         string  `json:"summary" binding:"required,min=10"`
	DeliveryAddress *string `json:"delivery_address,omitempty"`
}
type CreateOrderRequest struct {
	Summary            string                       `json:"summary" binding:"required,min=10"`
	DeliveryPreference domain.DeliveryPreference    `json:"delivery_preference" binding:"required,oneof=IN_STORE DELIVERY CURBSIDE"`
	DeliveryAddress    *string                      `json:"delivery_address,omitempty"`
	PostalCode         *string                      `json:"postal_code,omitempty"`
	SelectedProducts   *[]domain.AISuggestedProduct `json:"selected_products,omitempty"`
}
type GetOrdersFilters struct {
	Status             *string `form:"status"`
	DeliveryPreference *string `form:"delivery_preference"`
	SortBy             string  `form:"sort_by"`
	SortOrder          string  `form:"sort_order"`
	Limit              int     `form:"limit"`
	Offset             int     `form:"offset"`
}
type UpdateOrderRequest struct {
	Status              *domain.OrderStatus          `json:"status,omitempty"`
	AISuggestedProducts *[]domain.AISuggestedProduct `json:"ai_suggested_products,omitempty"`
}
type orderService struct {
	orderRepo repository.OrderRepository
	aiService AIService
}
func NewOrderService(orderRepo repository.OrderRepository, aiService AIService) OrderService {
	return &orderService{
		orderRepo: orderRepo,
		aiService: aiService,
	}
}
func (s *orderService) GetAISuggestions(req *GetAISuggestionsRequest) ([]domain.AISuggestedProduct, error) {
	if s.aiService == nil {
		return []domain.AISuggestedProduct{}, nil
	}
	products, err := s.aiService.SuggestProducts(req.Summary, req.DeliveryAddress)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *orderService) CreateOrder(userID uint, req *CreateOrderRequest) (*domain.Order, error) {
	if req.Summary == "" {
		return nil, ErrInvalidInput
	}
	validPreferences := []domain.DeliveryPreference{
		domain.DeliveryPreferenceInStore,
		domain.DeliveryPreferenceDelivery,
		domain.DeliveryPreferenceCurbside,
	}
	validPreference := false
	for _, vp := range validPreferences {
		if req.DeliveryPreference == vp {
			validPreference = true
			break
		}
	}
	if !validPreference {
		return nil, ErrInvalidInput
	}
	if req.DeliveryPreference == domain.DeliveryPreferenceDelivery {
		if req.DeliveryAddress == nil || *req.DeliveryAddress == "" {
			return nil, ErrInvalidInput
		}
	}
	order := &domain.Order{
		UserID:             userID,
		Summary:            req.Summary,
		DeliveryPreference: req.DeliveryPreference,
		DeliveryAddress:    req.DeliveryAddress,
		PostalCode:         req.PostalCode,
		Status:             domain.OrderStatusPending,
	}
	if req.SelectedProducts != nil && len(*req.SelectedProducts) > 0 {
		if err := order.SetAISuggestedProducts(*req.SelectedProducts); err != nil {
			return nil, ErrInvalidInput
		}
	}
	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}
func (s *orderService) GetOrders(userID uint, filters *GetOrdersFilters) ([]*domain.Order, error) {
	if filters == nil {
		filters = &GetOrdersFilters{}
	}
	if filters.Limit == 0 {
		filters.Limit = 100
	}
	repoFilters := repository.OrderFilters{
		Status:             filters.Status,
		DeliveryPreference: filters.DeliveryPreference,
		SortBy:             filters.SortBy,
		SortOrder:          filters.SortOrder,
		Limit:              filters.Limit,
		Offset:             filters.Offset,
	}
	return s.orderRepo.GetByUserIDWithFilters(userID, repoFilters)
}
func (s *orderService) GetOrderByID(orderID, userID uint) (*domain.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, ErrOrderNotFound
	}
	if order.UserID != userID {
		return nil, ErrUnauthorizedAccess
	}
	return order, nil
}
func (s *orderService) UpdateOrder(orderID, userID uint, req *UpdateOrderRequest) (*domain.Order, error) {
	order, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return nil, ErrOrderNotFound
	}
	if order.UserID != userID {
		return nil, ErrUnauthorizedAccess
	}
	if req.Status != nil {
		validStatuses := []domain.OrderStatus{
			domain.OrderStatusPending,
			domain.OrderStatusProcessing,
			domain.OrderStatusCompleted,
			domain.OrderStatusCancelled,
		}
		valid := false
		for _, vs := range validStatuses {
			if *req.Status == vs {
				valid = true
				break
			}
		}
		if !valid {
			return nil, ErrInvalidOrderStatus
		}
		order.Status = *req.Status
	}
	if req.AISuggestedProducts != nil {
		if err := order.SetAISuggestedProducts(*req.AISuggestedProducts); err != nil {
			return nil, ErrInvalidInput
		}
	}
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}
	return order, nil
}
