package ports

type TopologyRepository interface {
	Save(s string)
	Topologies() []string
}

type MessagesRepository interface {
	Save(id int)
	MessageExists(id int) bool
	Messages() map[int]struct{}
	MessagesCount() int
}
