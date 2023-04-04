package main

import (
	"log"
	"maelstrom-broadcast/handlers"
	"maelstrom-broadcast/repositories"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	id := uuid.New()
	n := maelstrom.NewNode()

	m := repositories.NewMessagesRepositoryInMemory()
	t := repositories.NewTopologyRepositoryInMemory()
	l := handlers.NewLogger(id.String())
	defer l.Close()

	n.Handle("broadcast", handlers.BroadCastHandlerFactory(l, n, m, t))
	n.Handle("read", handlers.ReadHandlerFactory(n, m))
	n.Handle("topology", handlers.TopologyHandlerFactory(n, t))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
