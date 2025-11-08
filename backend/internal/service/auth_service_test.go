package service_test
import (
	"testing"
	"weel-backend/internal/domain"
	"weel-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)
type MockUserRepositoryForAuth struct {
	mock.Mock
}
func (m *MockUserRepositoryForAuth) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepositoryForAuth) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserRepositoryForAuth) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserRepositoryForAuth) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepositoryForAuth) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockUserRepositoryForAuth) List(limit, offset int) ([]*domain.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}
func (m *MockUserRepositoryForAuth) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
type AuthServiceTestSuite struct {
	suite.Suite
	authService service.AuthService
	mockRepo    *MockUserRepositoryForAuth
	jwtService  *service.JWTService
}
func (suite *AuthServiceTestSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepositoryForAuth)
	suite.jwtService = service.NewJWTService()
	suite.authService = service.NewAuthService(suite.mockRepo, suite.jwtService)
}
func (suite *AuthServiceTestSuite) TestLogin_Success() {
	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := service.HashPassword(password)
	user := &domain.User{
		ID:       1,
		Email:    email,
		Password: hashedPassword,
	}
	suite.mockRepo.On("GetByEmail", email).Return(user, nil)
	suite.mockRepo.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)
	response, err := suite.authService.Login(email, password)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), response)
	assert.NotEmpty(suite.T(), response.Token)
	assert.Equal(suite.T(), user.ID, response.User.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *AuthServiceTestSuite) TestLogin_InvalidCredentials() {
	email := "test@example.com"
	password := "wrongpassword"
	hashedPassword, _ := service.HashPassword("correctpassword")
	user := &domain.User{
		ID:       1,
		Email:    email,
		Password: hashedPassword,
	}
	suite.mockRepo.On("GetByEmail", email).Return(user, nil)
	response, err := suite.authService.Login(email, password)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), response)
	assert.Equal(suite.T(), service.ErrInvalidCredentials, err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *AuthServiceTestSuite) TestGetCurrentUser_Success() {
	userID := uint(1)
	user := &domain.User{
		ID:    userID,
		Email: "test@example.com",
	}
	suite.mockRepo.On("GetByID", userID).Return(user, nil)
	result, err := suite.authService.GetCurrentUser(userID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, result.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}
func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
