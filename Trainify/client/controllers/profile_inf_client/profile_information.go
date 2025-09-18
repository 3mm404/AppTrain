package controllers_inf_client

import (
	dtos_user "github.com/3mm404/gymgo/client/dtos"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/models"
	"github.com/gofiber/fiber/v2"
)

func ShowProfileInformation(c *fiber.Ctx) error {
	// Obtener el ID del usuario del contexto
	userIDFloat := c.Locals("user").(float64)
	userID := uint(userIDFloat)

	var user models.User

	if err := config.MYDB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Usuario no encontrado",
		})
	}

	// Convertimos el modelo a DTO
	response := dtos_user.ShowInformationUserProfile{
		Nombre:          user.Nombre,
		Email:           user.Email,
		Telefono:        user.Telefono,
		FechaNacimiento: user.FechaNacimiento,
		Foto:            user.Foto,
	}

	return c.JSON(response)
}
