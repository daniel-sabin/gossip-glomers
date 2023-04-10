package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func TopologyHandlerFactory(l ports.Logger, n *maelstrom.Node, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		body, _ := parseBody(msg.Body)

		for _, v := range body["topology"].(map[string][]string)[n.ID()] {
			t.Save(v)
		}

		return n.Reply(msg, map[string]any{
			"type": "topology_ok",
		})
	}

}
