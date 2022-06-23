package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "register success",
	})
}
