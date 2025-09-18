package controllers_gym

import (
	dtos "github.com/3mm404/gymgo/client/dtos"
	userservices "github.com/3mm404/gymgo/client/services"
	"github.com/gofiber/fiber/v2"
)

// RegisterUserWithGym registra un usuario con un gimnasio
// el gimnacio lo registra sin contraseña
func RegisterUserWithGym(c *fiber.Ctx) error {
	var dto dtos.RegisterUserDTO

	// Parsear el JSON del cuerpo de la petición
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	// Registrar el usuario
	user, err := userservices.RegisterUser(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al registrar usuario: " + err.Error(),
		})
	}

	// No devolvemos la contraseña
	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario registrado exitosamente",
		"user":    user,
	})

}
