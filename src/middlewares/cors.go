package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (mc *MiddlewaresContainer) Cors(c *fiber.Ctx) error {
	corsConfig := cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	}
	corsSetUp := cors.New(corsConfig)
	c.App().Use(corsSetUp)
	return c.Next()
}
