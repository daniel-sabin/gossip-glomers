package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func BroadCastHandlerFactory(n *maelstrom.Node, m ports.MessagesRepository, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		body, _ := parseBody(msg.Body)

		id := int(body["message"].(float64))

		if m.MessageExists(id) {
			return nil
		}
		m.Save(id)

		for _, v := range t.Topologies() {
			// Skip to send message to Src client
			if msg.Src == v {
				continue
			}

			go func(nod string) {
				n.Send(nod, body)
			}(v)
		}

		// ACK Current message recieved
		return n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})
	}
}
