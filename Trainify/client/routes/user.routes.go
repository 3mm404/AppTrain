package routes_user

import (
	controllers_client "github.com/3mm404/gymgo/client/controllers"
	controllers_client_memberships "github.com/3mm404/gymgo/client/controllers/memberships"
	controllers_inf_client "github.com/3mm404/gymgo/client/controllers/profile_inf_client"

	"github.com/3mm404/gymgo/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(route *fiber.App) {

	// Registro de clientes
	route.Post("/users/register", controllers_client.RegisterUser)

	// Inicio de sesión para clientes
	route.Post("/users/login", controllers_client.LoginUser)

	// Informacion de Usuario Logueado
	route.Get("/users/profile", middleware.JWTProtected(), middleware.GetProfile)

	//Informacion de membresía del cliente
	route.Get("/users/membership", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_inf_client.GetMembershipStatus)

	// Listar todos los gimnasios
	route.Get("/users/gyms", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_client.ShowAllGyms)

	// Listar membresías disponibles por gimnasio
	route.Get("/users/:id/memberships", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_client_memberships.GetMembershipsByGymID)
	//por realizar
	// Consultar membresía activa del cliente
	route.Get("/memberships/active", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_client_memberships.GetActiveMembership)

	// Comprar una membresía
	route.Post("/memberships/purchase", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_client_memberships.PurchaseMembership)

	// Historial de pagos del cliente
	route.Get("/payments/history", middleware.JWTProtected(), middleware.RoleProtected("cliente"), controllers_client_memberships.GetPaymentHistory)
}
