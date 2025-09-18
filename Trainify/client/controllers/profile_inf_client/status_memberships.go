package controllers_inf_client

import (
	"time"

	dtos_user "github.com/3mm404/gymgo/client/dtos"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetMembershipStatus obtiene el estado de la membresía del usuario
func GetMembershipStatus(c *fiber.Ctx) error {
	// Obtener el ID del usuario desde el contexto JWT
	userID := c.Locals("user").(uint)

	// Buscar la información de la membresía del usuario
	var user models.User
	if err := config.MYDB.Preload("Gym").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener la información del usuario",
		})
	}

	// Si el usuario no tiene gym asignado, retornamos un DTO vacío
	if user.GymID == nil {
		return c.JSON(dtos_user.InfMembershipClient{
			GymID:     0,
			GymName:   "",
			Status:    "inactivo",
			StartDate: time.Time{},
			EndDate:   time.Time{},
		})
	}

	// Buscamos la membresía activa del usuario
	var membership models.UserMembership
	if err := config.MYDB.Where("user_id = ? AND activo = true", userID).First(&membership).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Si no hay membresía activa, retornamos un DTO con estado inactivo
			return c.JSON(dtos_user.InfMembershipClient{
				GymID:     *user.GymID,
				GymName:   "",
				Status:    "inactivo",
				StartDate: time.Time{},
				EndDate:   time.Time{},
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener la membresía del usuario",
		})
	}

	// Retornar la información de la membresía
	return c.JSON(dtos_user.InfMembershipClient{
		GymID:     *user.GymID,
		GymName:   "",
		Status:    "activo",
		StartDate: membership.FechaInicio,
		EndDate:   membership.FechaFin,
	})
}
