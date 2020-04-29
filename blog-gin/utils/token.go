/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-01-09 10:51
 */
package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**
 * 生成token
 */
func GenerateToken(name string, password string) string {
	hmacSampleSecret := []byte("my_secret_ky")
	//进行 算法加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": name,
		"password": password,
		"nbf":      time.Date(2020, 01, 04, 12, 0, 0, 0, time.UTC).Unix(),
	})
	if tokenString, err := token.SignedString(hmacSampleSecret); err != nil {

		return " "
	} else {

		return tokenString
	}

}

/**
 * 解析Token
 *
 */
func ParsingToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("my_secret_ky"), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err

}
