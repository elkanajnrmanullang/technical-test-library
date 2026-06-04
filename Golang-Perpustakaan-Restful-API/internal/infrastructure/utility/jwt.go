package utility

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTUtility struct {
	config config.JwtConfig
}

func NewJWTUtility(cfg config.JwtConfig) *JWTUtility {
	return &JWTUtility{
		config: cfg,
	}
}

// GenerateToken generates a new JWT token for a given user ID.
func (j *JWTUtility) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.config.GetJwtExpirationTime()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.Secret))
}

// GenerateRefreshToken generates a new refresh JWT token.
func (j *JWTUtility) GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(j.config.GetJwtRefreshExpirationTime()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.RefreshSecret))
}

// ParseToken parses and validates a JWT token.
func (j *JWTUtility) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

// ParseRefreshToken parses and validates a refresh JWT token.
func (j *JWTUtility) ParseRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.config.RefreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
