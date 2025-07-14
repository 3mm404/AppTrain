package main

import (
	"TrainApp/config"
	"TrainApp/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Conectar a la base de datos
	config.ConectionDB()
	//Migrar la base de datos
	database.Migrate()

	app := fiber.New()

	app.Listen(":8080")
}
