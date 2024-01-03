package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-api/handlers"
	"go-api/validations"
)

func SetupRouter(r *fiber.App) {
	r.Get("/", validations.HeaderValidator, handlers.Home)
	r.Post("/autentication", handlers.Authentication)
}
