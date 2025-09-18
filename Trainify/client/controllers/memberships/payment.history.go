package controllers_client_memberships

import (
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func GetPaymentHistory(c *fiber.Ctx) error {
	userID := c.Locals("user").(uint)

	var payments []models.Payment
	if result := config.MYDB.Where("user_id = ?", userID).Find(&payments); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener historial de pagos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"payments": payments,
	})
}
