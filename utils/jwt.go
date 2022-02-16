package utils

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// access secret key
var accessKey = ""

// Claims ...
type Claims struct {
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	jwt.StandardClaims
}

// ExtractJWTString Get claim from token string
func ExtractJWTString(tokenString string) (*Claims, error) {
	retClaim := &Claims{}
	JwtToken, err := jwt.ParseWithClaims(tokenString, retClaim, func(t *jwt.Token) (interface{}, error) {
		return []byte(accessKey), nil
	})
	if err == nil {
		if !JwtToken.Valid {
			return retClaim, nil
		}
	}
	return retClaim, err
}

// GenerateHash password hash generate
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword compare password and hash
func ComparePassword(password, hash string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash)); err != nil {
		return false, err
	}
	return true, nil
}
