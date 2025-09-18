package controllers_gym_memberships

import (
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

// DeleteMembershipType elimina un tipo de membresía existente de la base de datos
func DeleteMembershipType(c *fiber.Ctx) error {
	// Obtener el gymID del token
	gymID, ok := c.Locals("user").(uint) // Asumiendo que has almacenado el GymID como "user" en el token
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Gym ID no encontrado en el token"})
	}

	// Obtener el ID del tipo de membresía a eliminar desde los parámetros de la URL
	membershipTypeID := c.Params("id")
	if membershipTypeID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "ID de tipo de membresía no proporcionado"})
	}

	// Buscar el tipo de membresía en la base de datos
	var membershipType models.MembershipType
	if err := config.MYDB.Where("id = ? AND gym_id = ?", membershipTypeID, gymID).First(&membershipType).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Tipo de membresía no encontrado"})
	}

	// Eliminar el tipo de membresía de la base de datos
	if err := config.MYDB.Delete(&membershipType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error al eliminar el tipo de membresía"})
	}

	// Retornar una respuesta indicando que la eliminación fue exitosa
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Tipo de membresía eliminado con éxito"})
}
