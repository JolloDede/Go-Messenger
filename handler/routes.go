package handler

import (
	"github.com/JolloDede/Go-Messenger.git/cmd"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")

	app.Get("/home", HandleHome)
	app.All("/auth/register", handleRegister)
}

func HandleHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Fiber Template",
	}, "main")
}

func handleRegister(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		if cmd.KeyIsValid(c.FormValue("key")) {
			user := cmd.User{
				Username: c.FormValue("uname"),
				Password: c.FormValue("password"),
			}

			cmd.SaveUser(&user)

			return c.Redirect("/home")
		}
	}

	return c.Render("register", fiber.Map{}, "main")
}
