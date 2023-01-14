package scenes

import (
	"github.com/immersivesky/botsplugingo/updates"
)

func Message(update *updates.Update) updates.ReplyJSON {
	reply := updates.CreateReply()

	if update.Message.Text == "commands" {
		reply.AddBlock(updates.Action{
			Event: "message",
			Message: updates.Message{
				Chat: updates.Chat{
					ID: update.Message.Chat.ID,
				},
				Text: "Test commands:",
			},
		})
	}

	return reply.JSON()
}