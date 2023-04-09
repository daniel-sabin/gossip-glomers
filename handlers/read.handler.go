package handlers

import (
	"maelstrom-broadcast/ports"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func ReadHandlerFactory(n *maelstrom.Node, m ports.MessagesRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		time.Sleep(200 * time.Millisecond)
		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": m.Messages(),
		})
	}
}
