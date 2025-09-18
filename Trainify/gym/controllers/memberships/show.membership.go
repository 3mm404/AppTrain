package controllers_gym_memberships

import (
	"github.com/3mm404/gymgo/config"
	dtos_gym "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func ShowMembershipTypes(c *fiber.Ctx) error {
	// Lógica para mostrar los tipos de membresía disponibles
	// Obtener el gymID desde la URL
	gymID, ok := c.Locals("user").(uint) // Asumiendo que has almacenado el GymID como "user" en el token
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Gym ID no encontrado en el token"})
	}

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
			ID:           m.ID,         //identificador de membresia
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
