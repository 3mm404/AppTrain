package controllers_gym_memberships

import (
	"github.com/3mm404/gymgo/config"
	dtos_gym "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateMembershipType(c *fiber.Ctx) error {
	// Obtener el ID desde los parámetros de la URL
	id := c.Params("id")

	// Buscar el registro existente
	var membership models.MembershipType
	if err := config.MYDB.First(&membership, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Membresía no encontrada"})
	}

	// Parsear el JSON recibido en un DTO
	var updateDTO dtos_gym.MembershipTypeDTO
	if err := c.BodyParser(&updateDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Datos inválidos"})
	}

	// Validar el DTO
	if err := validate.Struct(updateDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	// Obtener el gymID desde el token (y verificar si tiene permiso para modificar)
	gymID, ok := c.Locals("user").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Gym ID no encontrado en el token"})
	}
	if membership.GymID != gymID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "No tienes permiso para modificar esta membresía"})
	}

	// Actualizar los campos
	membership.Nombre = updateDTO.Nombre
	membership.Descripcion = updateDTO.Descripcion
	membership.Precio = updateDTO.Precio
	membership.DuracionDias = updateDTO.DuracionDias
	membership.Status = updateDTO.Status

	// Guardar los cambios
	if err := config.MYDB.Save(&membership).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error al actualizar la membresía"})
	}

	return c.JSON(fiber.Map{"message": "Membresía actualizada exitosamente"})
}
