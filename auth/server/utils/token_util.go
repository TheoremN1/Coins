package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/TheoremN1/Coins/auth/envs"
	"github.com/TheoremN1/Coins/auth/server/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateTokens(userId string) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"role":    "user",                                // TODO: потом изменить
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Срок действия 24 часа
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 600).Unix(), // Срок действия месяц
	})

	signedAccessToken, err := accessToken.SignedString([]byte(envs.GetJwtSecret()))
	if err != nil {
		return models.Tokens{}, err
	}
	signedRefreshToken, err := refreshToken.SignedString([]byte(envs.GetJwtSecret()))
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

func ValidateRefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(envs.GetJwtSecret()), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]
		if userIdStr, ok := userId.(string); ok {
			return userIdStr, nil
		}
	}

	return "", fmt.Errorf("невозможно извлечь user_id из токена")
}

func ExtractDataFromAccessToken(tokenString string) (gin.H, error) {
	// Отсечение префикса "Bearer" из заголовка
	str := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))

	// Проверяем, что токен валиден
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный алгоритм подписи: %v", token.Header["alg"])
		}
		return []byte(envs.GetJwtSecret()), nil
	})
	if err != nil {
		return gin.H{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"]
		userRole := claims["role"]

		userIdStr, okId := userId.(string)
		userRoleStr, okRole := userRole.(string)
		if !(okId && okRole) {
			return gin.H{}, fmt.Errorf("невозможно извлечь данные")
		}
		return gin.H{
			"user_id":   userIdStr,
			"user_role": userRoleStr,
		}, nil
	}

	return gin.H{}, fmt.Errorf("недействительный токен")
}
