package pluginsgo

import (
	"github.com/immersivesky/botsplugingo/updates"
	"github.com/gofiber/fiber/v2"
)

type Plugin struct {
	Name	string				`json:"name,omitempty"`
	Scenes	map[string]Scene	`json:"scenes,omitempty"`
}

func Create(name string) *Plugin {
	return &Plugin{
		Name: name,
		Scenes: map[string]Scene{},
	}
}

type Scene func(*updates.Update) updates.ReplyJSON

func(plugin *Plugin) Set(event string, scene Scene) {
	if plugin.Scenes == nil {
		plugin.Scenes = map[string]Scene{}
	}
	plugin.Scenes[event] = scene
}

func(plugin *Plugin) Fiber(context *fiber.Ctx) error {
	update := updates.Update{}
	context.BodyParser(&update)

	scene := plugin.Scenes[update.Event]
	if scene == nil {
		return context.SendString(`{"event": "error", "error": {"about": "Scene not found"}}`)
	}

	return context.SendString(string(scene(&update)))
}