package controllers_client

import (
	dtos_user "github.com/3mm404/gymgo/client/dtos"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func ShowAllGyms(c *fiber.Ctx) error {

	var gyms []models.Gym
	var gymsDTO []dtos_user.ShowAllGymsDTO

	if err := config.MYDB.Find(&gyms).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "No se pudieron obtener los gimnasios",
		})
	}

	// Mapear los gimnasios a DTOs
	for _, gym := range gyms {
		gymsDTO = append(gymsDTO, dtos_user.ShowAllGymsDTO{
			ID:        gym.ID,
			Nombre:    gym.Nombre,
			Direccion: gym.Direccion,
			Telefono:  gym.Telefono,
			Foto:      gym.Foto,
			Latitud:   gym.Latitud,
			Longitud:  gym.Longitud,
		})
	}

	// Retornar los DTOs como respuesta en formato JSON
	return c.JSON(gymsDTO)
}
