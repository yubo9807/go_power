package service

import (
	"errors"
	"fmt"
	"server/src/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtType struct{}

var Jwt jwtType

var secret = "lifby9m2"
var secretMap = make(map[string]string) // 利用本地缓存充当 session

// 储存偷啃（在修改密码后重新执行，原先的偷啃将无效）
func (j *jwtType) StorageSetToken(key, tokenString string) {
	secretMap[key] = tokenString
}

// 获取缓存偷啃
func (j *jwtType) StorageGetToken(key string) string {
	return secretMap[key]
}

// 颁发 JWT
func (j *jwtType) Publish(info map[string]interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"info": info,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
		"iss":  "power-system",
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

// 验证 JWT
func (j *jwtType) Verify(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		info, _ := utils.InterfaceToMap(claims["info"])
		return info, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
