package service

import (
	"errors"
	"log"
	"time"
	"weel-backend/config"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
type JWTService struct {
	secretKey []byte
	expiresIn time.Duration
}

func NewJWTService() *JWTService {
	cfg, err := config.Load()
	if err != nil {
		log.Println("Warning: Failed to load config, using default JWT secret")
		return &JWTService{
			secretKey: []byte("default-secret-change-in-production"),
			expiresIn: 24 * time.Hour,
		}
	}
	secret := cfg.JWT.Secret
	if secret == "" {
		log.Println("Warning: JWT_SECRET not set, using default")
		secret = "default-secret-change-in-production"
	}
	return &JWTService{
		secretKey: []byte(secret),
		expiresIn: 24 * time.Hour,
	}
}
func (s *JWTService) GenerateToken(userID uint, email string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}
func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
