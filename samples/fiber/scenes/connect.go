package scenes

import (
	"github.com/immersivesky/botsplugingo/updates"
)

func Connect(update *updates.Update) updates.ReplyJSON {
	reply := updates.CreateReply()

	if update.Connect == "request" {
		reply.AddBlock(updates.Action{
			Event: "connect",
			Connect: "online",
		})
	}

	return reply.JSON()
}