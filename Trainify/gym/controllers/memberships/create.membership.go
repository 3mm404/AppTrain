package controllers_gym_memberships

import (
	"time"

	"github.com/3mm404/gymgo/config"
	dtos_gym "github.com/3mm404/gymgo/gym/dtos"
	"github.com/3mm404/gymgo/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// CrearMembershipType crea un nuevo tipo de membresía en la base de datos
func CrearMembershipType(c *fiber.Ctx) error {
	// Crear un objeto DTO para recibir los datos
	var membershipTypeDTO dtos_gym.MembershipTypeDTO

	// Parsear el JSON recibido en el DTO (sin el gym_id)
	if err := c.BodyParser(&membershipTypeDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Datos inválidos"})
	}

	// Validar el DTO
	if err := validate.Struct(membershipTypeDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	gymID, ok := c.Locals("user").(uint) // Asumiendo que has almacenado el GymID como "user" en el token
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Gym ID no encontrado en el token"})
	}

	// Ahora puedes usar gymID para crear la instancia de MembershipType
	membershipType := models.MembershipType{
		GymID:        gymID, // Usamos el gym_id extraído del JWT
		Nombre:       membershipTypeDTO.Nombre,
		Descripcion:  membershipTypeDTO.Descripcion,
		Precio:       membershipTypeDTO.Precio,
		DuracionDias: membershipTypeDTO.DuracionDias,
		Status:       membershipTypeDTO.Status,
		CreadoEn:     time.Now(),
	}

	// Guardar el nuevo tipo de membresía en la base de datos
	if err := config.MYDB.Create(&membershipType).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error al guardar el tipo de membresía"})
	}

	// Retornar la respuesta con los datos del tipo de membresía creado
	return c.Status(fiber.StatusCreated).JSON(membershipType)
}
