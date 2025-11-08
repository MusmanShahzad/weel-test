package service
import "errors"
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrOrderNotFound      = errors.New("order not found")
	ErrInvalidOrderStatus = errors.New("invalid order status")
	ErrUnauthorizedAccess = errors.New("unauthorized to access this order")
)
