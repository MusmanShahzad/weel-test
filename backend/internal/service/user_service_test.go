package service_test
import (
	"testing"
	"weel-backend/internal/domain"
	"weel-backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)
type MockUserRepository struct {
	mock.Mock
}
func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockUserRepository) List(limit, offset int) ([]*domain.User, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}
func (m *MockUserRepository) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
type UserServiceTestSuite struct {
	suite.Suite
	userService service.UserService
	mockRepo    *MockUserRepository
}
func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.userService = service.NewUserService(suite.mockRepo)
}
func (suite *UserServiceTestSuite) TestCreateUser_Success() {
	req := &service.CreateUserRequest{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.mockRepo.On("GetByEmail", "test@example.com").Return(nil, assert.AnError)
	suite.mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)
	user, err := suite.userService.CreateUser(req)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), user)
	assert.Equal(suite.T(), req.Email, user.Email)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *UserServiceTestSuite) TestCreateUser_EmailExists() {
	req := &service.CreateUserRequest{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	existingUser := &domain.User{ID: 1, Email: "test@example.com"}
	suite.mockRepo.On("GetByEmail", "test@example.com").Return(existingUser, nil)
	user, err := suite.userService.CreateUser(req)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), service.ErrEmailExists, err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *UserServiceTestSuite) TestGetUserByID_Success() {
	expectedUser := &domain.User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.mockRepo.On("GetByID", uint(1)).Return(expectedUser, nil)
	user, err := suite.userService.GetUserByID(1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser.ID, user.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *UserServiceTestSuite) TestGetUserByID_NotFound() {
	suite.mockRepo.On("GetByID", uint(1)).Return(nil, assert.AnError)
	user, err := suite.userService.GetUserByID(1)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), user)
	assert.Equal(suite.T(), service.ErrUserNotFound, err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *UserServiceTestSuite) TestUpdateUser_Success() {
	existingUser := &domain.User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	req := &service.UpdateUserRequest{
		FirstName: stringPtr("Updated"),
	}
	suite.mockRepo.On("GetByID", uint(1)).Return(existingUser, nil)
	suite.mockRepo.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)
	user, err := suite.userService.UpdateUser(1, req)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Updated", user.FirstName)
	suite.mockRepo.AssertExpectations(suite.T())
}
func (suite *UserServiceTestSuite) TestDeleteUser_Success() {
	existingUser := &domain.User{ID: 1}
	suite.mockRepo.On("GetByID", uint(1)).Return(existingUser, nil)
	suite.mockRepo.On("Delete", uint(1)).Return(nil)
	err := suite.userService.DeleteUser(1)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}
func stringPtr(s string) *string {
	return &s
}
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
