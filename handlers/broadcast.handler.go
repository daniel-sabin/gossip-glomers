package handlers

import (
	"context"
	"maelstrom-broadcast/ports"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func retry(n *maelstrom.Node, dest string, body map[string]any) error {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := n.SyncRPC(ctxTimeout, dest, body)
	return err
}

func BroadCastHandlerFactory(l ports.Logger, n *maelstrom.Node, m ports.MessagesRepository, t ports.TopologyRepository) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		body, err := parseBody(msg.Body)
		if err != nil {
			return nil
		}
		go n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})

		message := int(body["message"].(float64))
		if m.MessageExists(message) {
			return nil
		}
		m.Save(message)

		for _, dest := range t.Topologies() {
			// Skip to send message to Src client
			if msg.Src == dest {
				continue
			}

			go func(dest string) {
				if err := retry(n, dest, body); err != nil {
					for {
						time.Sleep(time.Second)
						if err := retry(n, dest, body); err == nil {
							break
						}
					}
				}
			}(dest)
		}
		return nil
	}
}
