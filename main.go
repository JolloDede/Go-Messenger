package main

import (
	"github.com/JolloDede/Go-Messenger.git/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	handler.Setup(app)
	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })

	app.Listen(":3000")
}
