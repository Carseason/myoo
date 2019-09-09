package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtKey []byte
)

func NewJwtKey(key string) {
	jwtKey = []byte(key)
}

/*******************JWT生成**********************/
func Encode(id int64) string {
	date := time.Now().Unix()
	if jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": date,           //处理时间,在此时间前不可用
		"iat": date,           //签发时间
		"exp": date + 1209600, //到期时间
		"iss": "Myoo",         //签发人,iss
		"id":  id,
	}).SignedString(jwtKey); err == nil {
		return jwtToken
	}
	return ""
}

/*******************JWT检索**********************/
func Decode(tokenString string) (int64, bool) {
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	}); err == nil {
		if value, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if v, ok := value["id"].(float64); ok {
				return int64(v), true
			}
		}
	}
	return 0, false
}

/*******************JWT Id字段读取**********************/
func Check_Id(tokenString string) int64 {
	if token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	}); err != nil {
		return 0
	} else {
		if value, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if v, ok := value["id"].(float64); ok {
				return int64(v)
			}
		}
	}
	return 0
}
