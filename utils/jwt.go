package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("graphqldemo")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(168 * time.Hour) // 有效期一周

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "demo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func ValidateJWT(token string) error {
	if token == "" {
		return errors.New("Authorization token must be present")
	}

	claims, err := ParseToken(token)
	if err != nil {
		return err
	} else if time.Now().Unix() > claims.ExpiresAt {
		return errors.New("Error. Token is expired")
	}

	return nil
}
