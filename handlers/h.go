package handlers

import (
	"github.com/gofiber/fiber/v2"

	"go-api/models"
)

// HandlerGetStatus docs
//
// @Summary  Get Status
// @Description xxxhh
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200		{string}  string  "OK"
// @Success		201		{string}  string  "Status Created"
// @Failure		400		{string}  error  "Bad Request"
// @Failure		500		{string}  error  "Internal server error"
// @Router			/ [get]
// @security ApiKeyAuth
func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"succes": true,
	})
}

// HandlerAutentication docs
//
// @Summary  Authentication
// @Description User Autentication
// @Tags			users
// @Accept			json
// @Produce		json
// @Param payload body models.Authentication true "Authentication"
// @Success		200		{string}  string  "OK"
// @Failure		400		{string}  error  "Bad Request"
// @Router			/autentication [post]
func Authentication(c *fiber.Ctx) error {
	reqModel := models.Authentication{}

	if err := c.BodyParser(&reqModel); err != nil {
		return c.JSON(err)
	}

	return c.JSON(reqModel)
}
