package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func ReadHandlerFactory(n *maelstrom.Node, m ports.MessagesRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": m.Messages(),
		})
	}
}
