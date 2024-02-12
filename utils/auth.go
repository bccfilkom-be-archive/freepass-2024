package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"freepass-2024/models"
)

func FindUserByEmail(email string) *models.User {
	for _, user := range models.Users{
		if user.Email == email {
			return &user
		}
	}
	return nil
}

func IsDuplicate(username, email string) bool {
	for _, user := range models.Users {
		if user.Name == username || user.Email == email {
			return true
		}
	}
	return false
}

func GenerateToken(ID int) (string, error) {
	index, user := FindUserByID(ID)

	if index == -1 {
		// Handle not found
		return "", nil
	}

	// Contoh pembuatan token JWT (gunakan library jwt-go untuk implementasi yang lebih kuat)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Token berlaku selama 1 jam
	})

	// Simpan token sebagai string dalam User
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func FindUserByID(userID int) (int, *models.User) {
	for i, user := range models.Users {
		if user.ID == userID {
			return i, &user
		}
	}
	return -1, nil
}

func GetUserIDFromToken(c *gin.Context) (int, error) {
	tokenHeader := c.GetHeader("Authorization")
	parts := strings.Split(tokenHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return 0, err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if id, ok := claims["id"].(float64); ok {
				index, user := FindUserByID(int(id))

				if index == -1 {
					return 0, nil
				}

				if user.Token != parts[1] {
					return 0, nil
				}

				return int(id), nil
			}
		}
	}
	return 0, nil
}

func GetUserRoleFromToken(c *gin.Context) string {
	// Implementasikan pengambilan peran pengguna dari token di sini
	// Contoh sederhana, ambil dari header Authorization
	tokenHeader := c.GetHeader("Authorization")
	parts := strings.Split(tokenHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		token, _ := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims["role"].(string)
		}
	}
	return ""
}
