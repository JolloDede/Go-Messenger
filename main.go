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

	app.Listen(":3000")
}
