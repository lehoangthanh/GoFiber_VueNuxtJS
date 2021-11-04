package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"time"
)

const SecretKey = "secret"
var ExpireTime = time.Now().Add(time.Minute * 5)

type ClaimsWScope struct {
	jwt.StandardClaims
	Scope string
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	
	payload := token.Claims.(*ClaimsWScope)
	isAmbassador := strings.Contains(c.Path(), "api/ambassador")
	
	if (payload.Scope == "admin" && isAmbassador) || (payload.Scope == "ambassador" && !isAmbassador) {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}
	
	return c.Next()
}

func GenerateJWT(id uint, scope string) (string, error){
	payload := ClaimsWScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Scope = scope
	
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

func GetUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return 0, err
	}
	
	payload := token.Claims.(*ClaimsWScope)
	id, _ := strconv.Atoi(payload.Subject)
	return uint(id), nil
}

func IsAmbassador(c *fiber.Ctx) bool {
	return strings.Contains(c.Path(), "api/ambassador")
}