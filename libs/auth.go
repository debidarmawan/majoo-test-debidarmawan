package libs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type TokenMetadata struct {
	Expires int64
	UserID  uint64
}

func GenerateToken(userID uint64) string {
	secretKey := os.Getenv("SECRET_KEY")
	expTime := time.Now().Add(time.Hour * 24).Unix()

	claims := jwt.MapClaims{}
	claims["exp"] = expTime
	claims["user_id"] = userID

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return token
}

func AuthValidator(c *fiber.Ctx) (*TokenMetadata, error) {
	var tokenString string

	bearerToken := c.Get("Authorization")
	tokenRaw := strings.Split(bearerToken, " ")
	if len(tokenRaw) == 2 {
		tokenString = tokenRaw[1]
	}
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expires := int64(claims["exp"].(float64))
		userID := uint64(claims["user_id"].(float64))

		return &TokenMetadata{
			Expires: expires,
			UserID:  userID,
		}, nil
	}

	return nil, err

}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("SECRET_KEY")), nil
}
