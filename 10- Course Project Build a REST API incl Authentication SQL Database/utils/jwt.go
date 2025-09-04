package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new JWT token for a user
func GenerateJWT(email string, userId int) (string, error) {
	expirationTime := time.Now().Add(2 * time.Hour) // token valid for 2 hours
	claims := &Claims{
		Email: email,
		ID:    uint(userId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey) // return the signed token
}

// ValidateJWT verifies if the token is valid and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Type assertion
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GetUserIDFromToken extracts the user ID directly from a JWT token
func GetUserIDFromToken(tokenString string) (uint, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.ID, nil
}
