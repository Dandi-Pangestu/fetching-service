package middleware

import (
	"efishery.com/micro/shared/domains"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type CustomClaims struct {
	jwt.StandardClaims
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	TimeStamp string `json:"timestamp"`
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
}

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("application.resources.jwt.secret")), nil
	})

	if err == nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, domains.ErrUnauthenticate
}
