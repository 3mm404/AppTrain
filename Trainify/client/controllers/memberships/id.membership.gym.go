//MOSTRAR TODAS LAS MEMBRESIAS DE UN GIMNASIO

package controllers_client_memberships

import (
	"github.com/3mm404/gymgo/config"
	dtos_gym "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func GetMembershipsByGymID(c *fiber.Ctx) error {
	// Obtener el gymID desde la URL
	gymID := c.Params("id")

	// Buscar todas las membresías con preload del gimnasio
	var memberships []models.MembershipType
	if err := config.MYDB.Preload("Gym").Where("gym_id = ?", gymID).Find(&memberships).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudieron obtener las membresías",
		})
	}

	// Mapear al DTO de respuesta
	var result []dtos_gym.ShowGymMoreMembershipDTO
	for _, m := range memberships {
		dto := dtos_gym.ShowGymMoreMembershipDTO{
			Gym:          m.Gym.Nombre, // Aquí accedes al nombre del gimnasio
			Descripcion:  m.Descripcion,
			Precio:       m.Precio,
			DuracionDias: m.DuracionDias,
			Status:       m.Status,
		}
		result = append(result, dto)
	}

	return c.JSON(result)
}
