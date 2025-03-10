package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")

	app.Get("/", HandleHome)
}

func HandleHome(c *fiber.Ctx) error {

	// home := views.HomeIndex(false)
	// handler := adaptor.HTTPHandler(templ.Handler(home))

	// return handler(c)

	return c.Render("index", fiber.Map{
		"Title": "Go Fiubert Template",
	}, "main")
}
