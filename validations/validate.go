package validations

import (
	"github.com/gofiber/fiber/v2"
)

const (
	DefaultApiKeyHeaderName = "x-api-key"
	DefaultApiKeyQueryName  = "key"
)

func HeaderValidator(c *fiber.Ctx) error {
	apikey := "test"
	// fmt.Println(apikey)
	// export API_KEY=xxx
	// fmt.Printf(c.Request().Header.String())

	header := c.Get(DefaultApiKeyHeaderName)
	query := c.Query(DefaultApiKeyQueryName)

	if header == "" && query == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "unauthorized",
		})
	}

	if header != "" && header == apikey {
		return c.Next()
	}

	if query != "" && query == apikey {
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "unauthorized",
	})
}
