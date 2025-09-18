package controllers_gym

import (
	dtos "github.com/3mm404/gymgo/gym/dtos" // asegúrate que esta ruta sea correcta
	gym_services "github.com/3mm404/gymgo/gym/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterGym(c *fiber.Ctx) error {
	var gymDTO dtos.RegisterGymDTO

	// Parsear el JSON del cuerpo de la petición
	if err := c.BodyParser(&gymDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "algo no esta bien: " + err.Error(),
		})
	}

	// Establecer valores predeterminados si es necesario
	gymDTO.SetDefaultValues()

	// Registrar el gimnasio con el servicio
	createdGym, err := gym_services.RegisterGym(gymDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al registrar gimnasio: " + err.Error(),
		})
	}

	// Devolvemos el gimnasio creado
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Gimnasio en espera de aprobación",
		"gym":     createdGym,
	})
}
