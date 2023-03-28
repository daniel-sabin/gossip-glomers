package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func ReadHandlerFactory(n *maelstrom.Node, m ports.MessagesRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		ids := make([]int, 0, m.MessagesCount())
		for id := range m.Messages() {
			ids = append(ids, id)
		}

		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": ids,
		})
	}
}
