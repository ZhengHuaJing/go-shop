package util

import (
	"github.com/google/uuid"
	"github.com/zhenghuajing/fresh_shop/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int
	UserName string
	PassWord string
	UUID     uuid.UUID
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id int, userName, password string, uuid uuid.UUID) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(9999 * time.Hour)

	claims := Claims{
		ID:       id,
		UserName: userName,
		PassWord: password,
		UUID:     uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "LotteryDraw",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(global.Config.MD5.JwtSecret))

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.MD5.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
