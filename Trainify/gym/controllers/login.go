package controllers_gym

import (
	"errors"
	"os"
	"time"

	"github.com/3mm404/gymgo/config"
	dtos_gym "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginGym(c *fiber.Ctx) error {
	input := dtos_gym.LoginGymDTO{}

	// Parsear datos del cuerpo de la solicitud
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	// Validar que el identificador no esté vacío
	if input.Identificador == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El identificador (email o teléfono) no puede estar vacío.",
		})
	}

	var user models.Gym

	// Buscar el usuario por email o teléfono usando una consulta más explícita
	result := config.MYDB.Table("gyms").
		Where("email = ? OR telefono = ?", input.Identificador, input.Identificador).
		Select("id", "nombre", "email", "telefono", "password").
		First(&user)

	// Si no se encuentra el usuario
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Usuario no encontrado",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al buscar el usuario: " + result.Error.Error(),
		})
	}

	// Verificar la contraseña con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Credenciales incorrectas",
		})
	}

	// Obtener la clave JWT de la variable de entorno
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "pq9ubrvuq43nv31nv91835vn98nv918nqv5938n8g" // Fallback para desarrollo
	}

	// Crear los claims del token
	claims := jwt.MapClaims{
		"user":     user.ID,
		"email":    user.Email,
		"telefono": user.Telefono,
		"rol":      "gym",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Crear y firmar el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "No se pudo generar el token: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Login exitoso",
		"name":     user.Nombre,
		"email":    user.Email,
		"telefono": user.Telefono,
		"token":    tokenString,
	})
}
