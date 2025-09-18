package controllers_client

import (
	dtos "github.com/3mm404/gymgo/client/dtos"
	userservices "github.com/3mm404/gymgo/client/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var dto dtos.RegisterUserDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	user, err := userservices.RegisterUser(dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := fiber.Map{
		"id":        user.ID,
		"nombre":    user.Nombre,
		"email":     user.Email,
		"telefono":  user.Telefono,
		"creado_en": user.CreadoEn,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario registrado exitosamente",
		"user":    response,
	})
}
