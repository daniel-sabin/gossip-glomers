package handlers

import (
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func TopologyHandlerFactory(n *maelstrom.Node, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		body, _ := parseBody(msg.Body)

		for k, v := range body["topology"].(map[string]any) {
			if k == n.ID() {
				for _, nodeName := range v.([]any) {
					t.Save(nodeName.(string))
				}
				break
			}
		}

		return n.Reply(msg, map[string]any{
			"type": "topology_ok",
		})
	}

}
