package routes_gym

import (
	controllers "github.com/3mm404/gymgo/gym/controllers"
	controllers_gym_memberships "github.com/3mm404/gymgo/gym/controllers/memberships"
	"github.com/3mm404/gymgo/middleware"
	"github.com/gofiber/fiber/v2"
)

func GymRoute(route *fiber.App) {

	// Registro de un nuevo gimnasio
	route.Post("/gym/register", controllers.RegisterGym)
	// Informacion de gimnasio Logueado
	route.Get("/gym/profile", middleware.JWTProtected(), middleware.RoleProtected("gym"), controllers.ShowProfileInformation)
	// Inicio de sesión del gimnasio
	route.Post("/gym/login", controllers.LoginGym)

	// Registro de usuarios del gimnasio
	route.Post("/gym/register/user", controllers.RegisterUserWithGym)

	// Dashboard privado del gimnasio (requiere autenticación)
	route.Get("/gym/dashboard", middleware.JWTProtected(), middleware.RoleProtected("gym"), func(c *fiber.Ctx) error {
		return c.SendString("Bienvenido al dashboard del gimnasio.")
	})
	//
	//
	//

	// Crear tipo de membresía (solo para gimnasios autenticados)
	route.Post("gym/crear/membresias", middleware.JWTProtected(), middleware.RoleProtected("gym"), controllers_gym_memberships.CrearMembershipType)
	// Listar membresías disponibles por gimnasio
	route.Get("/gym/mymemberships", middleware.JWTProtected(), middleware.RoleProtected("gym"), controllers_gym_memberships.ShowMembershipTypes)
	// Aquí puedes agregar:
	route.Put("gym/membership-types/:id", middleware.JWTProtected(), middleware.RoleProtected("gym"), controllers_gym_memberships.UpdateMembershipType)
	//Eliminar Menbresia
	route.Delete("gym/eliminar/membresias/:id", middleware.JWTProtected(), middleware.RoleProtected("gym"), controllers_gym_memberships.DeleteMembershipType)
	//x
	// - GET /gyms/users: Ver usuarios registrados en el gimnasio

	// - Otros endpoints relacionados al gimnasio
}
