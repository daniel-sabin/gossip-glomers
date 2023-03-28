package repositories

import "sync"

type TopologyRepositoryInMemory struct {
	store []string
	mutex sync.RWMutex
}

func NewTopologyRepositoryInMemory() *TopologyRepositoryInMemory {
	return &TopologyRepositoryInMemory{
		store: make([]string, 0),
	}
}

func (t *TopologyRepositoryInMemory) Save(s string) {
	t.mutex.Lock()
	t.store = append(t.store, s)
	t.mutex.Unlock()
}

func (t *TopologyRepositoryInMemory) Topologies() []string {
	return t.store
}
