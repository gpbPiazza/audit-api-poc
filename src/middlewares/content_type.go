package middlewares

import "github.com/gofiber/fiber/v2"

func (mc *MiddlewaresContainer) ContentType(c *fiber.Ctx) error {
	c.Type("json", "utf-8")

	return c.Next()
}
