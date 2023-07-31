package route

import (
	"github.com/gofiber/fiber/v2"
	"guying-go/internal/interceptor"
)

/**
 * @program: guying-go
 * @description:
 * @author: sickle
 * @create: 2023-07-27 10:17
 **/

func InitRouter(app *fiber.App) {

	api := app.Group("/api")
	//jwt查看https://github.com/gofiber/contrib/tree/main/jwt#hs256-test
	// JWT Middleware
	api.Use(interceptor.Auth())
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api.Get("/1", func(ctx *fiber.Ctx) error {
		return ctx.JSON("xxxx")
	})
}
