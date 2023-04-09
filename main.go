package main

import (
	"context"
	"log"
	"maelstrom-broadcast/handlers"
	"maelstrom-broadcast/repositories"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type FakeLogger struct {
}

func (F *FakeLogger) Debug(s string) {
}

func (F *FakeLogger) Close() {
}

func main() {
	n := maelstrom.NewNode()
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	m := repositories.NewMessagesRepositoryInMemory()
	t := repositories.NewTopologyRepositoryInMemory()
	l := handlers.NewLogger(n)
	defer l.Close()
	// l := &FakeLogger{}

	n.Handle("broadcast", handlers.BroadCastHandlerFactory(l, n, m, t, ctx))
	n.Handle("read", handlers.ReadHandlerFactory(n, m))
	n.Handle("topology", handlers.TopologyHandlerFactory(l, n, t))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
