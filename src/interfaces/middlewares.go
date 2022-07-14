package interfaces

import "github.com/gofiber/fiber/v2"

type MiddlewaresContainer interface {
	ContentType(c *fiber.Ctx) error
	Cors(c *fiber.Ctx) error
}