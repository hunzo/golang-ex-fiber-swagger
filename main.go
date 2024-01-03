package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "go-api/docs"
	"go-api/routes"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-api-key
// @description			Description for what is this security definition being used
func main() {
	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault)

	routes.SetupRouter(app)

	app.Listen(":8080")
}
