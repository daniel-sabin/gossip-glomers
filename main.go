package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Message struct {
	message string
}

func main() {

	n := maelstrom.NewNode()

	store := make(map[int]struct{})
	topology := make([]string, 0)
	var storeMutex sync.RWMutex

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		id := int(body["message"].(float64))
		storeMutex.Lock()
		if _, exists := store[id]; exists {
			storeMutex.Unlock()
			return nil
		}
		store[id] = struct{}{}
		storeMutex.Unlock()

		for _, v := range topology {
			// Skip to send message to Src client
			if msg.Src == v {
				continue
			}

			go func(nod string) {
				n.Send(nod, body)
			}(v)
		}
		return n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		storeMutex.RLock()
		ids := make([]int, 0, len(store))
		for id := range store {
			ids = append(ids, id)
		}
		storeMutex.RUnlock()

		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": ids,
		})
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		for k, v := range body["topology"].(map[string]any) {
			if k == n.ID() {
				for _, nodeName := range v.([]any) {
					topology = append(topology, nodeName.(string))
				}
				break
			}
		}

		return n.Reply(msg, map[string]any{
			"type": "topology_ok",
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
