package service_test

import (
	"testing"
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
	"weel-backend/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) GetByID(id uint) (*domain.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Order), args.Error(1)
}
func (m *MockOrderRepository) GetByUserID(userID uint, limit, offset int) ([]*domain.Order, error) {
	args := m.Called(userID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Order), args.Error(1)
}
func (m *MockOrderRepository) Update(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}
func (m *MockOrderRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockOrderRepository) GetByUserIDWithFilters(userID uint, filters repository.OrderFilters) ([]*domain.Order, error) {
	args := m.Called(userID, filters)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Order), args.Error(1)
}

type OrderServiceTestSuite struct {
	suite.Suite
	orderService service.OrderService
	mockRepo     *MockOrderRepository
}

func (suite *OrderServiceTestSuite) SetupTest() {
	suite.mockRepo = new(MockOrderRepository)
	suite.orderService = service.NewOrderService(suite.mockRepo, nil)
}
func (suite *OrderServiceTestSuite) TestCreateOrder_Delivery_Success() {
	userID := uint(1)
	postalCode := "12345"
	address := "123 Main St, New York, NY 10001"
	req := &service.CreateOrderRequest{
		Summary:            "I need groceries for the week including milk, bread, eggs, and vegetables",
		DeliveryPreference: domain.DeliveryPreferenceDelivery,
		DeliveryAddress:    &address,
		PostalCode:         &postalCode,
	}
	suite.mockRepo.On("Create", mock.AnythingOfType("*domain.Order")).Return(nil)
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), order)
	assert.Equal(suite.T(), userID, order.UserID)
	assert.Equal(suite.T(), domain.OrderStatusPending, order.Status)
	assert.Equal(suite.T(), req.Summary, order.Summary)
	assert.Equal(suite.T(), domain.DeliveryPreferenceDelivery, order.DeliveryPreference)
	assert.NotNil(suite.T(), order.DeliveryAddress)
	assert.Equal(suite.T(), address, *order.DeliveryAddress)
	assert.Equal(suite.T(), postalCode, *order.PostalCode)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestCreateOrder_InStore_Success() {
	userID := uint(1)
	req := &service.CreateOrderRequest{
		Summary:            "I need groceries for the week",
		DeliveryPreference: domain.DeliveryPreferenceInStore,
		DeliveryAddress:    nil,
		PostalCode:         nil,
	}
	suite.mockRepo.On("Create", mock.AnythingOfType("*domain.Order")).Return(nil)
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), order)
	assert.Equal(suite.T(), domain.DeliveryPreferenceInStore, order.DeliveryPreference)
	assert.Nil(suite.T(), order.DeliveryAddress)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestCreateOrder_Curbside_Success() {
	userID := uint(1)
	address := "123 Main St, New York, NY 10001"
	req := &service.CreateOrderRequest{
		Summary:            "I need groceries",
		DeliveryPreference: domain.DeliveryPreferenceCurbside,
		DeliveryAddress:    &address,
		PostalCode:         nil,
	}
	suite.mockRepo.On("Create", mock.AnythingOfType("*domain.Order")).Return(nil)
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), order)
	assert.Equal(suite.T(), domain.DeliveryPreferenceCurbside, order.DeliveryPreference)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestCreateOrder_Delivery_WithoutAddress() {
	userID := uint(1)
	req := &service.CreateOrderRequest{
		Summary:            "I need groceries",
		DeliveryPreference: domain.DeliveryPreferenceDelivery,
		DeliveryAddress:    nil,
	}
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), order)
	assert.Equal(suite.T(), service.ErrInvalidInput, err)
}
func (suite *OrderServiceTestSuite) TestCreateOrder_EmptySummary() {
	userID := uint(1)
	address := "123 Main St"
	req := &service.CreateOrderRequest{
		Summary:            "",
		DeliveryPreference: domain.DeliveryPreferenceDelivery,
		DeliveryAddress:    &address,
	}
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), order)
	assert.Equal(suite.T(), service.ErrInvalidInput, err)
}
func (suite *OrderServiceTestSuite) TestCreateOrder_InvalidPreference() {
	userID := uint(1)
	address := "123 Main St"
	req := &service.CreateOrderRequest{
		Summary:            "I need groceries",
		DeliveryPreference: domain.DeliveryPreference("INVALID"),
		DeliveryAddress:    &address,
	}
	order, err := suite.orderService.CreateOrder(userID, req)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), order)
	assert.Equal(suite.T(), service.ErrInvalidInput, err)
}
func (suite *OrderServiceTestSuite) TestGetOrderByID_Success() {
	orderID := uint(1)
	userID := uint(1)
	order := &domain.Order{
		ID:     orderID,
		UserID: userID,
	}
	suite.mockRepo.On("GetByID", orderID).Return(order, nil)
	result, err := suite.orderService.GetOrderByID(orderID, userID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), order.ID, result.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestGetOrderByID_Unauthorized() {
	orderID := uint(1)
	userID := uint(1)
	otherUserID := uint(2)
	order := &domain.Order{
		ID:     orderID,
		UserID: otherUserID,
	}
	suite.mockRepo.On("GetByID", orderID).Return(order, nil)
	result, err := suite.orderService.GetOrderByID(orderID, userID)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), service.ErrUnauthorizedAccess, err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestUpdateOrder_Success() {
	orderID := uint(1)
	userID := uint(1)
	status := domain.OrderStatusCompleted
	order := &domain.Order{
		ID:     orderID,
		UserID: userID,
		Status: domain.OrderStatusPending,
	}
	suite.mockRepo.On("GetByID", orderID).Return(order, nil)
	suite.mockRepo.On("Update", mock.AnythingOfType("*domain.Order")).Return(nil)
	req := &service.UpdateOrderRequest{
		Status: &status,
	}
	result, err := suite.orderService.UpdateOrder(orderID, userID, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), status, result.Status)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *OrderServiceTestSuite) TestUpdateOrder_InvalidStatus() {
	orderID := uint(1)
	userID := uint(1)
	invalidStatus := domain.OrderStatus("invalid")
	order := &domain.Order{
		ID:     orderID,
		UserID: userID,
	}
	suite.mockRepo.On("GetByID", orderID).Return(order, nil)
	req := &service.UpdateOrderRequest{
		Status: &invalidStatus,
	}
	result, err := suite.orderService.UpdateOrder(orderID, userID, req)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), service.ErrInvalidOrderStatus, err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}
