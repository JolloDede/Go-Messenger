package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")

	app.Get("/home", HandleHome)
	app.All("/auth/register", handleRegister)
}

func HandleHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Fiubert Template",
	}, "main")
}

type User struct {
	Id       int
	Username string
	password string
}

func handleRegister(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost {
		uname := c.FormValue("uname")
		fmt.Println(uname)
		return c.Redirect("/home")
	}

	return c.Render("register", fiber.Map{}, "main")
}
