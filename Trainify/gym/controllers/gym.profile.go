package controllers_gym

import (
	"github.com/gofiber/fiber/v2"
)

func ShowProfileInformation(c *fiber.Ctx) error {
	return c.SendString("Información del gimnasio")
}
