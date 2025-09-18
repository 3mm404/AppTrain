package main

import (
	"log"
	"os"

	routes_user "github.com/3mm404/gymgo/client/routes"
	"github.com/3mm404/gymgo/config"
	"github.com/3mm404/gymgo/database"

	routes_gym "github.com/3mm404/gymgo/gym/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Migración y conexión a la base de datos
	config.Connect()
	database.Migrate()

	// Inicializar la app de Fiber
	app := fiber.New()

	// Aquí va el logger y el CORS
	app.Use(logger.New()) // usa configuración por defecto o pasa config personalizada
	app.Use(cors.New())   // permitir acceso desde otros orígenes

	// Rutas para APIs de gym y user
	routes_gym.GymRoute(app)
	routes_user.UserRoute(app)

	// Puerto
	// Leer el puerto desde una variable de entorno o usar el 3000 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // valor por defecto
	}

	// Escuchar en el puerto configurado
	log.Fatal(app.Listen(":" + port))
}
