package middleware // Rutas exclusivas por tipo de rol

import (
	"os"
	"strings"

	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RoleProtected(expectedRole string) fiber.Handler {
	// Obtener la clave JWT de una variable de entorno
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		// Fallback a una clave por defecto solo para desarrollo
		jwtSecret = "pq9ubrvuq43nv31nv91835vn98nv918nqv5938n8g"
	}

	return func(c *fiber.Ctx) error {
		// Obtener el token de los encabezados de la solicitud
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Falta el token"})
		}

		// Eliminar la palabra "Bearer " del encabezado para obtener solo el token
		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parsear el token usando la clave secreta
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token inválido"})
		}

		// Extraer los claims del token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token inválido"})
		}

		// Obtener el rol del claim
		rol, ok := claims["rol"].(string)
		if !ok || rol != expectedRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Acceso denegado para el rol actual"})
		}

		// Obtener el ID del usuario y asegurar que es un float64 (puedes cambiarlo según lo que estés utilizando)
		userID, exists := claims["user"]
		if !exists {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user claim no encontrado"})
		}

		userIDFloat, ok := userID.(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user_id inválido"})
		}

		// Guardamos el ID de usuario y el rol en las variables locales para usarlas en la siguiente capa
		c.Locals("user", uint(userIDFloat))
		c.Locals("rol", rol)

		// Pasamos al siguiente middleware
		return c.Next()
	}
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user").(uint)

	var user models.User
	if err := config.MYDB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Usuario no encontrado",
		})
	}

	return c.JSON(fiber.Map{
		"id":        user.ID,
		"nombre":    user.Nombre,
		"email":     user.Email,
		"telefono":  user.Telefono,
		"tipo":      user.TipoUsuario,
		"creado_en": user.CreadoEn,
	})
}
