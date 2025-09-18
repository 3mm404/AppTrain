package controllers_client_memberships

import (
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

// Obtener la membresía activa
func GetActiveMembership(c *fiber.Ctx) error {
	userID := c.Locals("user").(uint) // Obtener el ID del usuario del token JWT

	var membership models.MembershipType
	if result := config.MYDB.Where("user_id = ? AND activo = ?", userID, true).First(&membership); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No tienes una membresía activa",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"membership": membership,
	})
}
