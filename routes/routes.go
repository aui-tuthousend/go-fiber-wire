package routes

import (
	"go-fiber-wire/container"
	"go-fiber-wire/internal/features/auth"
	"go-fiber-wire/internal/features/role"
	"go-fiber-wire/internal/features/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c *container.AppContainer) {
	api := app.Group("/api")

	role.RegisterRoute(api, c.RoleHandler)
	user.RegisterRoute(api, c.UserHandler)
	auth.RegisterRoute(api, c.AuthHandler)
}
