package main

import (
	"github.com/immersivesky/botsplugingo/samples/fiber/scenes"
	"github.com/gofiber/fiber/v2"
	plugins "github.com/immersivesky/botsplugingo"
)

func main() {
	var (
		app		= fiber.New()
		plugin	= plugins.Create("Test Plugin")
	)

	plugin.Set("connect", scenes.Connect)
	plugin.Set("message", scenes.Message)

	app.Post("/api/v1/test", plugin.Fiber)
	app.Listen(":3000")
}