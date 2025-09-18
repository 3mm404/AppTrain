package controllers_client_memberships

import (
	"time"

	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

// comprar una membresía
func PurchaseMembership(c *fiber.Ctx) error {
	userID := c.Locals("user").(uint)
	var input models.UserMembership
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	// Verificar que la membresía exista
	var membershipType models.MembershipType
	if result := config.MYDB.First(&membershipType, input.MembershipTypeID); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Membresía no encontrada",
		})
	}

	// Crear la membresía para el usuario
	newMembership := models.UserMembership{
		UserID:           userID,
		MembershipTypeID: input.MembershipTypeID,
		FechaInicio:      time.Now(),
		FechaFin:         time.Now().AddDate(0, 0, 30), // Ejemplo: membresía por 30 días
		Activo:           true,
	}
	if result := config.MYDB.Create(&newMembership); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al comprar la membresía",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Membresía comprada exitosamente",
		"membership": newMembership,
	})
}
