package handlers

import (
	"fmt"
	"maelstrom-broadcast/ports"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func TopologyHandlerFactory(l ports.Logger, n *maelstrom.Node, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		body, _ := parseBody(msg.Body)
		l.Debug(fmt.Sprint(body["topology"]))

		for k, v := range body["topology"].(map[string]any) {
			if k == n.ID() {
				for _, nodeName := range v.([]any) {
					l.Debug(fmt.Sprintf("topologie : %s => %s \n", n.ID(), nodeName.(string)))
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
