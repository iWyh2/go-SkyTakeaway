package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	errs "go-SkyTakeaway/common/errors"
	"go-SkyTakeaway/config"
	"time"
)

// GenerateJWT 生成JWT令牌
func GenerateJWT(claimName string, claimData string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		claimName: claimData,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	// 获得签名后的完整token
	signedToken, err := token.SignedString([]byte(config.ServerConfig.Jwt.SecretKey))
	return signedToken, err
}

// ParseJWT 解析JWT令牌
func ParseJWT(tokenStr string, claimName string) (string, error) {
	// 解析
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.ServerConfig.Jwt.SecretKey), nil
	})
	// 错误处理
	if err != nil {
		return "", err
	}
	// 获得携带的信息
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claimData, ok := claims[claimName].(string)
		if !ok {
			// 获得信息失败
			return "", errs.JwtParseError
		}
		// 解析成功
		return claimData, nil
	}
	// 此时解析失败
	return "", errs.JwtParseError
}
