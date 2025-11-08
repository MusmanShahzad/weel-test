package repository_test
import (
	"fmt"
	"testing"
	"weel-backend/internal/domain"
	"weel-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type UserRepositoryTestSuite struct {
	suite.Suite
	db           *gorm.DB
	userRepo     repository.UserRepository
	testUser     *domain.User
}
func (suite *UserRepositoryTestSuite) SetupSuite() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=weel_test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		suite.T().Skip("Skipping test: database not available")
		return
	}
	suite.db = db
	suite.userRepo = repository.NewUserRepository(db)
	suite.db.AutoMigrate(&domain.User{})
}
func (suite *UserRepositoryTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Migrator().DropTable(&domain.User{})
	}
}
func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
}
func (suite *UserRepositoryTestSuite) TestCreateUser() {
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	err := suite.userRepo.Create(user)
	assert.NoError(suite.T(), err)
	assert.NotZero(suite.T(), user.ID)
}
func (suite *UserRepositoryTestSuite) TestGetByID() {
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.userRepo.Create(user)
	found, err := suite.userRepo.GetByID(user.ID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, found.Email)
}
func (suite *UserRepositoryTestSuite) TestGetByEmail() {
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.userRepo.Create(user)
	found, err := suite.userRepo.GetByEmail("test@example.com")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, found.ID)
}
func (suite *UserRepositoryTestSuite) TestUpdateUser() {
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.userRepo.Create(user)
	user.FirstName = "Updated"
	err := suite.userRepo.Update(user)
	assert.NoError(suite.T(), err)
	updated, _ := suite.userRepo.GetByID(user.ID)
	assert.Equal(suite.T(), "Updated", updated.FirstName)
}
func (suite *UserRepositoryTestSuite) TestDeleteUser() {
	user := &domain.User{
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}
	suite.userRepo.Create(user)
	err := suite.userRepo.Delete(user.ID)
	assert.NoError(suite.T(), err)
	_, err = suite.userRepo.GetByID(user.ID)
	assert.Error(suite.T(), err)
}
func (suite *UserRepositoryTestSuite) TestListUsers() {
	for i := 0; i < 5; i++ {
		user := &domain.User{
			Email:     fmt.Sprintf("test%d@example.com", i),
			FirstName: "Test",
			LastName:  "User",
		}
		suite.userRepo.Create(user)
	}
	users, err := suite.userRepo.List(10, 0)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), users, 5)
}
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
