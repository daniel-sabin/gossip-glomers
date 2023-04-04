package handlers

import (
	"fmt"
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func BroadCastHandlerFactory(l *Logger, n *maelstrom.Node, m ports.MessagesRepository, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		l.Debug(fmt.Sprintf("node %s \n", n.ID()))

		body, _ := parseBody(msg.Body)

		id := int(body["message"].(float64))

		if m.MessageExists(id) {
			l.Debug(fmt.Sprintf("exists %s -> ? (id = %d)\n", msg.Src, id))
			return nil
		}
		m.Save(id)

		for _, dest := range t.Topologies() {
			// Skip to send message to Src client
			if msg.Src == dest {
				l.Debug(fmt.Sprintf("skipped %s -> %s (id = %d)\n", msg.Src, dest, id))
				continue
			}

			go func(dest string) {
				l.Debug(fmt.Sprintf("sent %s -> %s (id = %d)\n", msg.Src, dest, id))
				n.Send(dest, body)
			}(dest)
		}

		// ACK Current message recieved
		return n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})
	}
}
