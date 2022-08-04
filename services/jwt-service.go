package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IJWTService interface {
	GenerateToken(username string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService() IJWTService {
	return &JWTService{
		secretKey: getSecretKey(),
		issuer:    "davidkarelh",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secretkey"
	}
	return secret
}

func (jwtS *JWTService) GenerateToken(username string) string {
	claims := &JwtCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    jwtS.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtS.secretKey))

	if err != nil {
		panic(err)
	}
	return t
}

func (jwtS *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(jwtS.secretKey), nil
	})
}
