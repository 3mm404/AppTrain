package middleware // Acceso general autenticado

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "pq9ubrvuq43nv31nv91835vn98nv918nqv5938n8g"
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Falta el token"})
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token inválido: " + err.Error()})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token inválido"})
		}

		// ✅ Aquí cambiamos a "user"
		userID, ok := claims["user"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user_id inválido"})
		}

		c.Locals("user", uint(userID))

		if email, ok := claims["email"].(string); ok {
			c.Locals("email", email)
		}

		if rol, ok := claims["rol"].(string); ok {
			c.Locals("rol", rol)
		}

		return c.Next()
	}
}
