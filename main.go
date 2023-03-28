package main

import (
	"log"
	"maelstrom-broadcast/handlers"
	"maelstrom-broadcast/repositories"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	n := maelstrom.NewNode()

	m := repositories.NewMessagesRepositoryInMemory()
	t := repositories.NewTopologyRepositoryInMemory()

	n.Handle("broadcast", handlers.BroadCastHandlerFactory(n, m, t))
	n.Handle("read", handlers.ReadHandlerFactory(n, m))
	n.Handle("topology", handlers.TopologyHandlerFactory(n, t))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
