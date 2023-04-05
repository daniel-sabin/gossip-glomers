package handlers

import (
	"fmt"
	"maelstrom-broadcast/ports"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func BroadCastHandlerFactory(l *Logger, n *maelstrom.Node, m ports.MessagesRepository, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		l.Debug(fmt.Sprintf("node %s \n", n.ID()))

		body, _ := parseBody(msg.Body)
		message := int(body["message"].(float64))

		l.Debug(fmt.Sprintf("message = %d\n", message))

		if m.MessageExists(message) {
			// l.Debug(fmt.Sprintf("exists %s -> ? (id = %d)\n", msg.Src, id))
			return nil
		}
		m.Save(message)

		broadcastMsgIds := make(map[string]struct{})
		for _, dest := range t.Topologies() {
			// Skip to send message to Src client
			if msg.Src == dest {
				// l.Debug(fmt.Sprintf("skipped %s -> %s (id = %d)\n", msg.Src, dest, id))
				continue
			}
			// Register broadcast destinations
			broadcastMsgIds[dest] = struct{}{}

			handler := func(msg maelstrom.Message) error {
				delete(broadcastMsgIds, dest)
				return nil
			}

			go func(dest string) {
				for len(broadcastMsgIds) > 0 {
					n.RPC(dest, body, handler)
					time.Sleep(500 * time.Microsecond)
				}
			}(dest)
		}

		return n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})
	}
}
