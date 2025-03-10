package handler

import (
	"github.com/JolloDede/Go-Messenger.git/views"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Setup(app *fiber.App) {
	app.Get("/", HandleHome)
}

func HandleHome(c *fiber.Ctx) error {

	home := views.HomeIndex(false)
	handler := adaptor.HTTPHandler(templ.Handler(home))

	return handler(c)
}
