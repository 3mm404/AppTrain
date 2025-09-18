package controllers_client

import (
	"time"

	dtos_user "github.com/3mm404/gymgo/client/dtos"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(c *fiber.Ctx) error {
	input := dtos_user.LoginUserDTO{}

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

	user := models.User{}

	// Buscar el usuario por email o teléfono
	result := config.MYDB.Where("email = ? OR telefono = ?", input.Identificador, input.Identificador).First(&user)

	// Si no se encuentra el usuario
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Usuario no encontrado",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al buscar el usuario",
		})
	}

	// Verificar la contraseña con bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Credenciales incorrectas",
		})
	}

	// Crear los claims del token
	claims := jwt.MapClaims{
		"user":     user.ID,
		"email":    user.Email,
		"telefono": user.Telefono,
		"rol":      "cliente", // Asignamos el rol cliente
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Firmar el token con la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("pq9ubrvuq43nv31nv91835vn98nv918nqv5938n8g"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "No se pudo generar el token",
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
